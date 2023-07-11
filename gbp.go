package gbp

import (
	"errors"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/datastore"
	"github.com/fatih/color"

	postgresadapter "github.com/LucasToledoPereira/go-boilerplate/adapters/datastore/postgres"
	s3adapter "github.com/LucasToledoPereira/go-boilerplate/adapters/filestore/aws"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/filestore"
	"github.com/LucasToledoPereira/go-boilerplate/config"
	"github.com/LucasToledoPereira/go-boilerplate/internal/audit"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"

	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
	"github.com/LucasToledoPereira/go-boilerplate/internal/server"
)

/*
* Para a conexão com a blockchain e correto funcionamento nós vamos necessitar:
* * Contrato Deployer, que será responsável por fazer o deploy dos contratos ERC20, ERC721 e ERC1155
* * Adapter para Keystore, que será responsável por retornar as chaves públicas e privadas da empresa e dos players(usuários)
* * Adapter para blockchain que deverá receber qual a a chain a ser utilizada (provider) e efetuar o deploy, mint, transfer e burn dos tokens
* * * Nos iremos fornecer as ferramentas necessárias para a implementação do adapter utilizando EVM
* * * O contrato deployer deverá ser inicializado (feito deploy) por fora da plataforma e deverá ser configurado para utilização na mesma.
* * * Devemos fornecer as funções para fazer o deploy do contrato Deployer
 */
type AdapterFunc func() (adapter any, err error)
type ModuleFunc func(input *ModuleFuncInput) (err error)
type ModulesFuncChain []ModuleFunc
type ModuleFuncInput struct {
	Router    *router.Router
	Datastore datastore.IDatastoreAdapter
	Filestore filestore.IFilestoreAdapter
}

type Builder struct {
	modulesChain ModulesFuncChain
	datastore    datastore.IDatastoreAdapter
	filestore    filestore.IFilestoreAdapter
	//keystoreadapter
	//emailadapter
	router *router.Router
	//prehandler
	//posthandler
}

func postgresAdapterFunc() AdapterFunc {
	return func() (any, error) {
		return &postgresadapter.PostgresAdapter{}, nil
	}
}

func s3StorageAdapterFunc() AdapterFunc {
	return func() (any, error) {
		return &s3adapter.S3Adapter{}, nil
	}
}

/*
This is a function named "New" which returns a pointer to a new instance of the "Builder" struct.
The purpose of this function is to create a new instance of the builder object.
*/
func New() *Builder {
	return &Builder{}
}

/*
The Default function returns a new instance of the Builder struct with Postgres and S3 storage adapters already added.
This is a convenient way to get started with building an application without having to manually add the adapters.
*/
func Default() *Builder {
	bdl := New()
	bdl.Use(postgresAdapterFunc(), s3StorageAdapterFunc())
	return bdl
}

/*
The "Use" function takes in a variable number of AdapterFuncs and sets the corresponding adapters for the Builder object.
It checks if the adapter implements the IDatastoreAdapter or IFilestoreAdapter interface and sets the corresponding adapter accordingly.
*/
func (builder *Builder) Use(adapterFns ...AdapterFunc) {
	for _, adapterFn := range adapterFns {
		adapter, _ := adapterFn()
		if adp, ok := adapter.(datastore.IDatastoreAdapter); ok {
			builder.datastore = adp
		}
		if adp, ok := adapter.(filestore.IFilestoreAdapter); ok {
			builder.filestore = adp
		}
	}

}

/*
This function returns the datastore adapter that is being used by the builder.
If the adapter has already been set, it will be returned.
Otherwise, an error will be returned indicating that no adapter has been found.
*/
func (builder *Builder) Datastore() (datastore.IDatastoreAdapter, error) {
	if builder.datastore != nil {
		return builder.datastore, nil
	}
	return nil, errors.New(codes.DatastoreAdapterNotFound.String())
}

func (builder *Builder) Register(moduleFns ...ModuleFunc) {
	builder.modulesChain = append(builder.modulesChain, moduleFns...)
}

/*
This function returns the filestore adapter that is being used by the builder.
If the adapter has already been set, it will be returned.
Otherwise, an error will be returned indicating that no adapter has been found.
*/
func (builder *Builder) Filestore() (filestore.IFilestoreAdapter, error) {
	if builder.filestore != nil {
		return builder.filestore, nil
	}
	return nil, errors.New(codes.FilestoreAdapterNotFound.String())
}

/*
This function loads the API configuration from a "config.yml" file and prints a success message indicating that environment variables have been loaded successfully.
*/
func (builder *Builder) config() {
	config.ReadConfig(".")
	color.Green("Environment variables loaded successfully...")
}

/*
This is a method that connects to a database using a given datastore adapter.
It first initializes the connection and then runs a default migration.
If there are any errors during the connection or migration process, it returns an error.
Once the connection and migration are successful, it returns nil.
The method also prints out messages indicating the success of the connection and migration.
*/
func (builder *Builder) connect(ds datastore.IDatastoreAdapter) (err error) {
	//Initiate the connection with the database
	err = ds.New()
	if err != nil {
		return err
	}
	color.Green("Database connected successfully...")

	//Run default migration
	err = ds.Migrate()
	if err != nil {
		return err
	}
	color.Green("Database migrated successfully...")

	return nil
}

/*
It iterates through a list of module functions and calls each one with a "ModuleFuncInput" argument that contains references to the router, datastore, and filestore.
The function collects any errors returned by the module functions and returns them as a slice.
*/
func (builder *Builder) loadModules() (errs []error) {
	for _, moduleFn := range builder.modulesChain {
		err := moduleFn(&ModuleFuncInput{
			Router:    builder.router,
			Datastore: builder.datastore,
			Filestore: builder.filestore,
		})
		errs = append(errs, err)
	}
	return
}

/*
This function is responsible for initializing and configuring various components of the application such as the datastore, filestore, router, and modules.
First, the function calls the "config" method to load application settings.
Then, it initializes the datastore using the "Datastore" method and connects to it using the "connect" method.
If the audit setting is enabled, it creates a new audit instance using the "New" method of the "audit" package.
Next, the function initializes the filestore using the "Filestore" method and creates a new instance of it using the "New" method.
It then creates a new router using the "New" method of the "router" package.
Finally, the function loads all registered modules using the "loadModules" method and returns a new instance of the server using the "server.New" method.
Overall, this function is responsible for setting up the application's core components and preparing it for use.
*/
func (builder *Builder) Bootstrap() (sv *server.Server, err error) {
	builder.config()
	ds, err := builder.Datastore()
	if err != nil {
		panic(err.Error())
	}
	builder.connect(ds)

	if config.C.Settings.Audit {
		audit.New(ds)
	}

	fs, err := builder.Filestore()
	if err != nil {
		panic(err.Error())
	}
	fs.New()
	color.Green("Filestore created successfully...")

	//Criar rotas utilizando gin
	builder.router = router.New()
	color.Green("Router created successfully...")

	//Load registered modules
	builder.loadModules()
	color.Green("Modules loaded successfully...")
	return server.New(builder.router), nil
}

import { ethers } from "hardhat";

async function main() {
  const ContractFactory = await ethers.getContractFactory("ERC721Immutable");

  const instance = await ContractFactory.deploy("My Collection", "MC");
  await instance.deployed();

  console.log(`Contract deployed to ${instance.address}`);
}


async function deployer() {
  const ContractFactory = await ethers.getContractFactory("Deployer");
  let admin = ethers.Wallet.createRandom().address;

  const instance = await ContractFactory.deploy(admin);
  await instance.deployed();

  const erca = await instance.newImmutableERC721("My Collection", "MC")
  console.log(`Contract ERC721 deployed to ${erca.data}`);
  console.log(`Contract deployed to ${instance.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
deployer().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

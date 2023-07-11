import { expect } from "chai";
import { ethers } from "hardhat";

describe("ERC721Immutable", function () {
  it("Test contract", async function () {
    const ContractFactory = await ethers.getContractFactory("ERC721Immutable");

    const instance = await ContractFactory.deploy("Lucas' First Collection", "LFC");
    await instance.deployed();

    expect(await instance.name()).to.equal("Lucas' First Collection");
  });
});

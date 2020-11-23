const Migrations = artifacts.require("Migrations");

module.exports = function (deployer, net) {
  if (net == "deployment") {
    deployer.deploy(Migrations);
  }
};

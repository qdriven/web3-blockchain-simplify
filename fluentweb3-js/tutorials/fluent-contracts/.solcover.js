module.exports = {
    istanbulReporter: ["html", "lcov"],
    providerOptions: {
      mnemonic: process.env.MNEMONIC,
    },
    skipFiles: ['test','test/fuzzing/KeepersCounterEchidnaTest.sol', 'test/LinkToken.sol', 'test/MockOracle.sol', 'test/MockV3Aggregator.sol', 'test/VRFCoordinatorV2Mock.sol'],
};


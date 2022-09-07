package io.alpha.compound.contracts;

import io.reactivex.Flowable;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.Callable;

import org.web3j.abi.EventEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.DynamicArray;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.abi.datatypes.generated.Uint224;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.abi.datatypes.generated.Uint32;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteFunctionCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.BaseEventResponse;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tuples.generated.Tuple2;
import org.web3j.tuples.generated.Tuple3;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 5.0.0.
 */
@SuppressWarnings("rawtypes")
public class ComptrollerContract extends Contract {
    public static final String BINARY = "Bin file was not provided";

    public static final String FUNC_pendingAdmin = "pendingAdmin";

    public static final String FUNC__SETPENDINGADMIN = "_setPendingAdmin";

    public static final String FUNC_comptrollerImplementation = "comptrollerImplementation";

    public static final String FUNC__ACCEPTIMPLEMENTATION = "_acceptImplementation";

    public static final String FUNC_pendingComptrollerImplementation = "pendingComptrollerImplementation";

    public static final String FUNC__SETPENDINGIMPLEMENTATION = "_setPendingImplementation";

    public static final String FUNC__ACCEPTADMIN = "_acceptAdmin";

    public static final String FUNC_admin = "admin";

    public static final String FUNC__ADDCOMPMARKETS = "_addCompMarkets";

    public static final String FUNC__BECOME = "_become";

    public static final String FUNC__BORROWGUARDIANPAUSED = "_borrowGuardianPaused";

    public static final String FUNC__DROPCOMPMARKET = "_dropCompMarket";

    public static final String FUNC__MINTGUARDIANPAUSED = "_mintGuardianPaused";

    public static final String FUNC__SETBORROWCAPGUARDIAN = "_setBorrowCapGuardian";

    public static final String FUNC__SETBORROWPAUSED = "_setBorrowPaused";

    public static final String FUNC__SETCLOSEFACTOR = "_setCloseFactor";

    public static final String FUNC__SETCOLLATERALFACTOR = "_setCollateralFactor";

    public static final String FUNC__SETCOMPRATE = "_setCompRate";

    public static final String FUNC__SETLIQUIDATIONINCENTIVE = "_setLiquidationIncentive";

    public static final String FUNC__SETMARKETBORROWCAPS = "_setMarketBorrowCaps";

    public static final String FUNC__SETMAXASSETS = "_setMaxAssets";

    public static final String FUNC__SETMINTPAUSED = "_setMintPaused";

    public static final String FUNC__SETPAUSEGUARDIAN = "_setPauseGuardian";

    public static final String FUNC__SETPRICEORACLE = "_setPriceOracle";

    public static final String FUNC__SETSEIZEPAUSED = "_setSeizePaused";

    public static final String FUNC__SETTRANSFERPAUSED = "_setTransferPaused";

    public static final String FUNC__SUPPORTMARKET = "_supportMarket";

    public static final String FUNC_ACCOUNTASSETS = "accountAssets";

    public static final String FUNC_ALLMARKETS = "allMarkets";

    public static final String FUNC_BORROWALLOWED = "borrowAllowed";

    public static final String FUNC_BORROWCAPGUARDIAN = "borrowCapGuardian";

    public static final String FUNC_BORROWCAPS = "borrowCaps";

    public static final String FUNC_BORROWGUARDIANPAUSED = "borrowGuardianPaused";

    public static final String FUNC_BORROWVERIFY = "borrowVerify";

    public static final String FUNC_CHECKMEMBERSHIP = "checkMembership";

    public static final String FUNC_claimComp = "claimComp";

    public static final String FUNC_CLOSEFACTORMANTISSA = "closeFactorMantissa";

    public static final String FUNC_COMPACCRUED = "compAccrued";

    public static final String FUNC_COMPBORROWSTATE = "compBorrowState";

    public static final String FUNC_COMPBORROWERINDEX = "compBorrowerIndex";

    public static final String FUNC_COMPCLAIMTHRESHOLD = "compClaimThreshold";

    public static final String FUNC_COMPINITIALINDEX = "compInitialIndex";

    public static final String FUNC_COMPRATE = "compRate";

    public static final String FUNC_COMPSPEEDS = "compSpeeds";

    public static final String FUNC_COMPSUPPLIERINDEX = "compSupplierIndex";

    public static final String FUNC_COMPSUPPLYSTATE = "compSupplyState";

    public static final String FUNC_ENTERMARKETS = "enterMarkets";

    public static final String FUNC_EXITMARKET = "exitMarket";

    public static final String FUNC_GETACCOUNTLIQUIDITY = "getAccountLiquidity";

    public static final String FUNC_GETALLMARKETS = "getAllMarkets";

    public static final String FUNC_GETASSETSIN = "getAssetsIn";

    public static final String FUNC_GETBLOCKNUMBER = "getBlockNumber";

    public static final String FUNC_GETCOMPADDRESS = "getCompAddress";

    public static final String FUNC_GETHYPOTHETICALACCOUNTLIQUIDITY = "getHypotheticalAccountLiquidity";

    public static final String FUNC_ISCOMPTROLLER = "isComptroller";

    public static final String FUNC_LIQUIDATEBORROWALLOWED = "liquidateBorrowAllowed";

    public static final String FUNC_LIQUIDATEBORROWVERIFY = "liquidateBorrowVerify";

    public static final String FUNC_LIQUIDATECALCULATESEIZETOKENS = "liquidateCalculateSeizeTokens";

    public static final String FUNC_LIQUIDATIONINCENTIVEMANTISSA = "liquidationIncentiveMantissa";

    public static final String FUNC_MARKETS = "markets";

    public static final String FUNC_MAXASSETS = "maxAssets";

    public static final String FUNC_MINTALLOWED = "mintAllowed";

    public static final String FUNC_MINTGUARDIANPAUSED = "mintGuardianPaused";

    public static final String FUNC_MINTVERIFY = "mintVerify";

    public static final String FUNC_ORACLE = "oracle";

    public static final String FUNC_PAUSEGUARDIAN = "pauseGuardian";

    public static final String FUNC_REDEEMALLOWED = "redeemAllowed";

    public static final String FUNC_REDEEMVERIFY = "redeemVerify";

    public static final String FUNC_REFRESHCOMPSPEEDS = "refreshCompSpeeds";

    public static final String FUNC_REPAYBORROWALLOWED = "repayBorrowAllowed";

    public static final String FUNC_REPAYBORROWVERIFY = "repayBorrowVerify";

    public static final String FUNC_SEIZEALLOWED = "seizeAllowed";

    public static final String FUNC_SEIZEGUARDIANPAUSED = "seizeGuardianPaused";

    public static final String FUNC_SEIZEVERIFY = "seizeVerify";

    public static final String FUNC_TRANSFERALLOWED = "transferAllowed";

    public static final String FUNC_TRANSFERGUARDIANPAUSED = "transferGuardianPaused";

    public static final String FUNC_TRANSFERVERIFY = "transferVerify";

    public static final Event NEWPENDINGIMPLEMENTATION_EVENT = new Event("NewPendingImplementation",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event NEWIMPLEMENTATION_EVENT = new Event("NewImplementation",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event NEWPENDINGADMIN_EVENT = new Event("NewPendingAdmin",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event NEWADMIN_EVENT = new Event("NewAdmin",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event FAILURE_EVENT = new Event("Failure",
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event ACTIONPAUSED_EVENT = new Event("ActionPaused",
            Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {
            }, new TypeReference<Bool>() {
            }));
    ;

    public static final Event COMPSPEEDUPDATED_EVENT = new Event("CompSpeedUpdated",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event DISTRIBUTEDBORROWERCOMP_EVENT = new Event("DistributedBorrowerComp",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {
            }, new TypeReference<Address>(true) {
            }, new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event DISTRIBUTEDSUPPLIERCOMP_EVENT = new Event("DistributedSupplierComp",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {
            }, new TypeReference<Address>(true) {
            }, new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event MARKETCOMPED_EVENT = new Event("MarketComped",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Bool>() {
            }));
    ;

    public static final Event MARKETENTERED_EVENT = new Event("MarketEntered",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event MARKETEXITED_EVENT = new Event("MarketExited",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event MARKETLISTED_EVENT = new Event("MarketListed",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }));
    ;

    public static final Event NEWBORROWCAP_EVENT = new Event("NewBorrowCap",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event NEWBORROWCAPGUARDIAN_EVENT = new Event("NewBorrowCapGuardian",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event NEWCLOSEFACTOR_EVENT = new Event("NewCloseFactor",
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event NEWCOLLATERALFACTOR_EVENT = new Event("NewCollateralFactor",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event NEWCOMPRATE_EVENT = new Event("NewCompRate",
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event NEWLIQUIDATIONINCENTIVE_EVENT = new Event("NewLiquidationIncentive",
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event NEWMAXASSETS_EVENT = new Event("NewMaxAssets",
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
            }, new TypeReference<Uint256>() {
            }));
    ;

    public static final Event NEWPAUSEGUARDIAN_EVENT = new Event("NewPauseGuardian",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    ;

    public static final Event NEWPRICEORACLE_EVENT = new Event("NewPriceOracle",
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
            }, new TypeReference<Address>() {
            }));
    private EthFilter filter;
    ;

    @Deprecated
    protected ComptrollerContract(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected ComptrollerContract(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected ComptrollerContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected ComptrollerContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }


    public RemoteFunctionCall<TransactionReceipt> _setPendingAdmin(String newPendingAdmin) {
        final Function function = new Function(
                FUNC__SETPENDINGADMIN,
                Arrays.<Type>asList(new Address(160, newPendingAdmin)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _acceptImplementation() {
        final Function function = new Function(
                FUNC__ACCEPTIMPLEMENTATION,
                Arrays.<Type>asList(),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setPendingImplementation(String newPendingImplementation) {
        final Function function = new Function(
                FUNC__SETPENDINGIMPLEMENTATION,
                Arrays.<Type>asList(new Address(160, newPendingImplementation)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _acceptAdmin() {
        final Function function = new Function(
                FUNC__ACCEPTADMIN,
                Arrays.<Type>asList(),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public List<NewPendingImplementationEventResponse> getNewPendingImplementationEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWPENDINGIMPLEMENTATION_EVENT, transactionReceipt);
        ArrayList<NewPendingImplementationEventResponse> responses = new ArrayList<NewPendingImplementationEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewPendingImplementationEventResponse typedResponse = new NewPendingImplementationEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldPendingImplementation = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newPendingImplementation = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewPendingImplementationEventResponse> newPendingImplementationEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewPendingImplementationEventResponse>() {
            @Override
            public NewPendingImplementationEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWPENDINGIMPLEMENTATION_EVENT, log);
                NewPendingImplementationEventResponse typedResponse = new NewPendingImplementationEventResponse();
                typedResponse.log = log;
                typedResponse.oldPendingImplementation = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newPendingImplementation = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewPendingImplementationEventResponse> newPendingImplementationEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWPENDINGIMPLEMENTATION_EVENT));
        return newPendingImplementationEventFlowable(filter);
    }

    public List<NewImplementationEventResponse> getNewImplementationEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWIMPLEMENTATION_EVENT, transactionReceipt);
        ArrayList<NewImplementationEventResponse> responses = new ArrayList<NewImplementationEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewImplementationEventResponse typedResponse = new NewImplementationEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldImplementation = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newImplementation = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewImplementationEventResponse> newImplementationEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewImplementationEventResponse>() {
            @Override
            public NewImplementationEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWIMPLEMENTATION_EVENT, log);
                NewImplementationEventResponse typedResponse = new NewImplementationEventResponse();
                typedResponse.log = log;
                typedResponse.oldImplementation = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newImplementation = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewImplementationEventResponse> newImplementationEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWIMPLEMENTATION_EVENT));
        return newImplementationEventFlowable(filter);
    }

    public List<NewPendingAdminEventResponse> getNewPendingAdminEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWPENDINGADMIN_EVENT, transactionReceipt);
        ArrayList<NewPendingAdminEventResponse> responses = new ArrayList<NewPendingAdminEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewPendingAdminEventResponse typedResponse = new NewPendingAdminEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldPendingAdmin = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newPendingAdmin = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewPendingAdminEventResponse> newPendingAdminEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewPendingAdminEventResponse>() {
            @Override
            public NewPendingAdminEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWPENDINGADMIN_EVENT, log);
                NewPendingAdminEventResponse typedResponse = new NewPendingAdminEventResponse();
                typedResponse.log = log;
                typedResponse.oldPendingAdmin = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newPendingAdmin = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewPendingAdminEventResponse> newPendingAdminEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWPENDINGADMIN_EVENT));
        return newPendingAdminEventFlowable(filter);
    }

    public List<NewAdminEventResponse> getNewAdminEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWADMIN_EVENT, transactionReceipt);
        ArrayList<NewAdminEventResponse> responses = new ArrayList<NewAdminEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewAdminEventResponse typedResponse = new NewAdminEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldAdmin = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newAdmin = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewAdminEventResponse> newAdminEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewAdminEventResponse>() {
            @Override
            public NewAdminEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWADMIN_EVENT, log);
                NewAdminEventResponse typedResponse = new NewAdminEventResponse();
                typedResponse.log = log;
                typedResponse.oldAdmin = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newAdmin = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewAdminEventResponse> newAdminEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWADMIN_EVENT));
        return newAdminEventFlowable(filter);
    }


    public List<ActionPausedEventResponse> getActionPausedEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(ACTIONPAUSED_EVENT, transactionReceipt);
        ArrayList<ActionPausedEventResponse> responses = new ArrayList<ActionPausedEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            ActionPausedEventResponse typedResponse = new ActionPausedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.action = (String) eventValues.getNonIndexedValues().get(1).getValue();
            typedResponse.pauseState = (Boolean) eventValues.getNonIndexedValues().get(2).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<ActionPausedEventResponse> actionPausedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, ActionPausedEventResponse>() {
            @Override
            public ActionPausedEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(ACTIONPAUSED_EVENT, log);
                ActionPausedEventResponse typedResponse = new ActionPausedEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.action = (String) eventValues.getNonIndexedValues().get(1).getValue();
                typedResponse.pauseState = (Boolean) eventValues.getNonIndexedValues().get(2).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<ActionPausedEventResponse> actionPausedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(ACTIONPAUSED_EVENT));
        return actionPausedEventFlowable(filter);
    }

    public List<CompSpeedUpdatedEventResponse> getCompSpeedUpdatedEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(COMPSPEEDUPDATED_EVENT, transactionReceipt);
        ArrayList<CompSpeedUpdatedEventResponse> responses = new ArrayList<CompSpeedUpdatedEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            CompSpeedUpdatedEventResponse typedResponse = new CompSpeedUpdatedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.newSpeed = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<CompSpeedUpdatedEventResponse> compSpeedUpdatedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, CompSpeedUpdatedEventResponse>() {
            @Override
            public CompSpeedUpdatedEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(COMPSPEEDUPDATED_EVENT, log);
                CompSpeedUpdatedEventResponse typedResponse = new CompSpeedUpdatedEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.newSpeed = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<CompSpeedUpdatedEventResponse> compSpeedUpdatedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(COMPSPEEDUPDATED_EVENT));
        return compSpeedUpdatedEventFlowable(filter);
    }

    public List<DistributedBorrowerCompEventResponse> getDistributedBorrowerCompEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(DISTRIBUTEDBORROWERCOMP_EVENT, transactionReceipt);
        ArrayList<DistributedBorrowerCompEventResponse> responses = new ArrayList<DistributedBorrowerCompEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            DistributedBorrowerCompEventResponse typedResponse = new DistributedBorrowerCompEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.borrower = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse.compDelta = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.compBorrowIndex = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<DistributedBorrowerCompEventResponse> distributedBorrowerCompEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, DistributedBorrowerCompEventResponse>() {
            @Override
            public DistributedBorrowerCompEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(DISTRIBUTEDBORROWERCOMP_EVENT, log);
                DistributedBorrowerCompEventResponse typedResponse = new DistributedBorrowerCompEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.borrower = (String) eventValues.getIndexedValues().get(1).getValue();
                typedResponse.compDelta = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.compBorrowIndex = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<DistributedBorrowerCompEventResponse> distributedBorrowerCompEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(DISTRIBUTEDBORROWERCOMP_EVENT));
        return distributedBorrowerCompEventFlowable(filter);
    }

    public List<DistributedSupplierCompEventResponse> getDistributedSupplierCompEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(DISTRIBUTEDSUPPLIERCOMP_EVENT, transactionReceipt);
        ArrayList<DistributedSupplierCompEventResponse> responses = new ArrayList<DistributedSupplierCompEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            DistributedSupplierCompEventResponse typedResponse = new DistributedSupplierCompEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.supplier = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse.compDelta = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.compSupplyIndex = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<DistributedSupplierCompEventResponse> distributedSupplierCompEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, DistributedSupplierCompEventResponse>() {
            @Override
            public DistributedSupplierCompEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(DISTRIBUTEDSUPPLIERCOMP_EVENT, log);
                DistributedSupplierCompEventResponse typedResponse = new DistributedSupplierCompEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.supplier = (String) eventValues.getIndexedValues().get(1).getValue();
                typedResponse.compDelta = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.compSupplyIndex = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<DistributedSupplierCompEventResponse> distributedSupplierCompEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(DISTRIBUTEDSUPPLIERCOMP_EVENT));
        return distributedSupplierCompEventFlowable(filter);
    }

    public List<FailureEventResponse> getFailureEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(FAILURE_EVENT, transactionReceipt);
        ArrayList<FailureEventResponse> responses = new ArrayList<FailureEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            FailureEventResponse typedResponse = new FailureEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.error = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.info = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            typedResponse.detail = (BigInteger) eventValues.getNonIndexedValues().get(2).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<FailureEventResponse> failureEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, FailureEventResponse>() {
            @Override
            public FailureEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(FAILURE_EVENT, log);
                FailureEventResponse typedResponse = new FailureEventResponse();
                typedResponse.log = log;
                typedResponse.error = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.info = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                typedResponse.detail = (BigInteger) eventValues.getNonIndexedValues().get(2).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<FailureEventResponse> failureEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(FAILURE_EVENT));
        return failureEventFlowable(filter);
    }

    public List<MarketCompedEventResponse> getMarketCompedEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(MARKETCOMPED_EVENT, transactionReceipt);
        ArrayList<MarketCompedEventResponse> responses = new ArrayList<MarketCompedEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            MarketCompedEventResponse typedResponse = new MarketCompedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.isComped = (Boolean) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<MarketCompedEventResponse> marketCompedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, MarketCompedEventResponse>() {
            @Override
            public MarketCompedEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(MARKETCOMPED_EVENT, log);
                MarketCompedEventResponse typedResponse = new MarketCompedEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.isComped = (Boolean) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<MarketCompedEventResponse> marketCompedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(MARKETCOMPED_EVENT));
        return marketCompedEventFlowable(filter);
    }

    public List<MarketEnteredEventResponse> getMarketEnteredEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(MARKETENTERED_EVENT, transactionReceipt);
        ArrayList<MarketEnteredEventResponse> responses = new ArrayList<MarketEnteredEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            MarketEnteredEventResponse typedResponse = new MarketEnteredEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.account = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<MarketEnteredEventResponse> marketEnteredEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, MarketEnteredEventResponse>() {
            @Override
            public MarketEnteredEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(MARKETENTERED_EVENT, log);
                MarketEnteredEventResponse typedResponse = new MarketEnteredEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.account = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<MarketEnteredEventResponse> marketEnteredEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(MARKETENTERED_EVENT));
        return marketEnteredEventFlowable(filter);
    }

    public List<MarketExitedEventResponse> getMarketExitedEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(MARKETEXITED_EVENT, transactionReceipt);
        ArrayList<MarketExitedEventResponse> responses = new ArrayList<MarketExitedEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            MarketExitedEventResponse typedResponse = new MarketExitedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.account = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<MarketExitedEventResponse> marketExitedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, MarketExitedEventResponse>() {
            @Override
            public MarketExitedEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(MARKETEXITED_EVENT, log);
                MarketExitedEventResponse typedResponse = new MarketExitedEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.account = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<MarketExitedEventResponse> marketExitedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(MARKETEXITED_EVENT));
        return marketExitedEventFlowable(filter);
    }

    public List<MarketListedEventResponse> getMarketListedEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(MARKETLISTED_EVENT, transactionReceipt);
        ArrayList<MarketListedEventResponse> responses = new ArrayList<MarketListedEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            MarketListedEventResponse typedResponse = new MarketListedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<MarketListedEventResponse> marketListedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, MarketListedEventResponse>() {
            @Override
            public MarketListedEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(MARKETLISTED_EVENT, log);
                MarketListedEventResponse typedResponse = new MarketListedEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<MarketListedEventResponse> marketListedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(MARKETLISTED_EVENT));
        return marketListedEventFlowable(filter);
    }

    public List<NewBorrowCapEventResponse> getNewBorrowCapEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWBORROWCAP_EVENT, transactionReceipt);
        ArrayList<NewBorrowCapEventResponse> responses = new ArrayList<NewBorrowCapEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewBorrowCapEventResponse typedResponse = new NewBorrowCapEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.newBorrowCap = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewBorrowCapEventResponse> newBorrowCapEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewBorrowCapEventResponse>() {
            @Override
            public NewBorrowCapEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWBORROWCAP_EVENT, log);
                NewBorrowCapEventResponse typedResponse = new NewBorrowCapEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.newBorrowCap = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewBorrowCapEventResponse> newBorrowCapEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWBORROWCAP_EVENT));
        return newBorrowCapEventFlowable(filter);
    }

    public List<NewBorrowCapGuardianEventResponse> getNewBorrowCapGuardianEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWBORROWCAPGUARDIAN_EVENT, transactionReceipt);
        ArrayList<NewBorrowCapGuardianEventResponse> responses = new ArrayList<NewBorrowCapGuardianEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewBorrowCapGuardianEventResponse typedResponse = new NewBorrowCapGuardianEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldBorrowCapGuardian = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newBorrowCapGuardian = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewBorrowCapGuardianEventResponse> newBorrowCapGuardianEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewBorrowCapGuardianEventResponse>() {
            @Override
            public NewBorrowCapGuardianEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWBORROWCAPGUARDIAN_EVENT, log);
                NewBorrowCapGuardianEventResponse typedResponse = new NewBorrowCapGuardianEventResponse();
                typedResponse.log = log;
                typedResponse.oldBorrowCapGuardian = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newBorrowCapGuardian = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewBorrowCapGuardianEventResponse> newBorrowCapGuardianEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWBORROWCAPGUARDIAN_EVENT));
        return newBorrowCapGuardianEventFlowable(filter);
    }

    public List<NewCloseFactorEventResponse> getNewCloseFactorEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWCLOSEFACTOR_EVENT, transactionReceipt);
        ArrayList<NewCloseFactorEventResponse> responses = new ArrayList<NewCloseFactorEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewCloseFactorEventResponse typedResponse = new NewCloseFactorEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldCloseFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newCloseFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewCloseFactorEventResponse> newCloseFactorEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewCloseFactorEventResponse>() {
            @Override
            public NewCloseFactorEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWCLOSEFACTOR_EVENT, log);
                NewCloseFactorEventResponse typedResponse = new NewCloseFactorEventResponse();
                typedResponse.log = log;
                typedResponse.oldCloseFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newCloseFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewCloseFactorEventResponse> newCloseFactorEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWCLOSEFACTOR_EVENT));
        return newCloseFactorEventFlowable(filter);
    }

    public List<NewCollateralFactorEventResponse> getNewCollateralFactorEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWCOLLATERALFACTOR_EVENT, transactionReceipt);
        ArrayList<NewCollateralFactorEventResponse> responses = new ArrayList<NewCollateralFactorEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewCollateralFactorEventResponse typedResponse = new NewCollateralFactorEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.oldCollateralFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            typedResponse.newCollateralFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(2).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewCollateralFactorEventResponse> newCollateralFactorEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewCollateralFactorEventResponse>() {
            @Override
            public NewCollateralFactorEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWCOLLATERALFACTOR_EVENT, log);
                NewCollateralFactorEventResponse typedResponse = new NewCollateralFactorEventResponse();
                typedResponse.log = log;
                typedResponse.cToken = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.oldCollateralFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                typedResponse.newCollateralFactorMantissa = (BigInteger) eventValues.getNonIndexedValues().get(2).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewCollateralFactorEventResponse> newCollateralFactorEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWCOLLATERALFACTOR_EVENT));
        return newCollateralFactorEventFlowable(filter);
    }

    public List<NewCompRateEventResponse> getNewCompRateEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWCOMPRATE_EVENT, transactionReceipt);
        ArrayList<NewCompRateEventResponse> responses = new ArrayList<NewCompRateEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewCompRateEventResponse typedResponse = new NewCompRateEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldCompRate = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newCompRate = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewCompRateEventResponse> newCompRateEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewCompRateEventResponse>() {
            @Override
            public NewCompRateEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWCOMPRATE_EVENT, log);
                NewCompRateEventResponse typedResponse = new NewCompRateEventResponse();
                typedResponse.log = log;
                typedResponse.oldCompRate = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newCompRate = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewCompRateEventResponse> newCompRateEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWCOMPRATE_EVENT));
        return newCompRateEventFlowable(filter);
    }

    public List<NewLiquidationIncentiveEventResponse> getNewLiquidationIncentiveEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWLIQUIDATIONINCENTIVE_EVENT, transactionReceipt);
        ArrayList<NewLiquidationIncentiveEventResponse> responses = new ArrayList<NewLiquidationIncentiveEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewLiquidationIncentiveEventResponse typedResponse = new NewLiquidationIncentiveEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldLiquidationIncentiveMantissa = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newLiquidationIncentiveMantissa = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewLiquidationIncentiveEventResponse> newLiquidationIncentiveEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewLiquidationIncentiveEventResponse>() {
            @Override
            public NewLiquidationIncentiveEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWLIQUIDATIONINCENTIVE_EVENT, log);
                NewLiquidationIncentiveEventResponse typedResponse = new NewLiquidationIncentiveEventResponse();
                typedResponse.log = log;
                typedResponse.oldLiquidationIncentiveMantissa = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newLiquidationIncentiveMantissa = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewLiquidationIncentiveEventResponse> newLiquidationIncentiveEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWLIQUIDATIONINCENTIVE_EVENT));
        return newLiquidationIncentiveEventFlowable(filter);
    }

    public List<NewMaxAssetsEventResponse> getNewMaxAssetsEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWMAXASSETS_EVENT, transactionReceipt);
        ArrayList<NewMaxAssetsEventResponse> responses = new ArrayList<NewMaxAssetsEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewMaxAssetsEventResponse typedResponse = new NewMaxAssetsEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldMaxAssets = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newMaxAssets = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewMaxAssetsEventResponse> newMaxAssetsEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewMaxAssetsEventResponse>() {
            @Override
            public NewMaxAssetsEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWMAXASSETS_EVENT, log);
                NewMaxAssetsEventResponse typedResponse = new NewMaxAssetsEventResponse();
                typedResponse.log = log;
                typedResponse.oldMaxAssets = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newMaxAssets = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewMaxAssetsEventResponse> newMaxAssetsEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWMAXASSETS_EVENT));
        return newMaxAssetsEventFlowable(filter);
    }

    public List<NewPauseGuardianEventResponse> getNewPauseGuardianEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWPAUSEGUARDIAN_EVENT, transactionReceipt);
        ArrayList<NewPauseGuardianEventResponse> responses = new ArrayList<NewPauseGuardianEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewPauseGuardianEventResponse typedResponse = new NewPauseGuardianEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldPauseGuardian = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newPauseGuardian = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewPauseGuardianEventResponse> newPauseGuardianEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewPauseGuardianEventResponse>() {
            @Override
            public NewPauseGuardianEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWPAUSEGUARDIAN_EVENT, log);
                NewPauseGuardianEventResponse typedResponse = new NewPauseGuardianEventResponse();
                typedResponse.log = log;
                typedResponse.oldPauseGuardian = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newPauseGuardian = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewPauseGuardianEventResponse> newPauseGuardianEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWPAUSEGUARDIAN_EVENT));
        return newPauseGuardianEventFlowable(filter);
    }

    public List<NewPriceOracleEventResponse> getNewPriceOracleEvents(TransactionReceipt transactionReceipt) {
        List<EventValuesWithLog> valueList = extractEventParametersWithLog(NEWPRICEORACLE_EVENT, transactionReceipt);
        ArrayList<NewPriceOracleEventResponse> responses = new ArrayList<NewPriceOracleEventResponse>(valueList.size());
        for (EventValuesWithLog eventValues : valueList) {
            NewPriceOracleEventResponse typedResponse = new NewPriceOracleEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.oldPriceOracle = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.newPriceOracle = (String) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<NewPriceOracleEventResponse> newPriceOracleEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, NewPriceOracleEventResponse>() {
            @Override
            public NewPriceOracleEventResponse apply(Log log) {
                EventValuesWithLog eventValues = extractEventParametersWithLog(NEWPRICEORACLE_EVENT, log);
                NewPriceOracleEventResponse typedResponse = new NewPriceOracleEventResponse();
                typedResponse.log = log;
                typedResponse.oldPriceOracle = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse.newPriceOracle = (String) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<NewPriceOracleEventResponse> newPriceOracleEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(NEWPRICEORACLE_EVENT));
        return newPriceOracleEventFlowable(filter);
    }

    public RemoteFunctionCall<TransactionReceipt> _addCompMarkets(List<String> cTokens) {
        final Function function = new Function(
                FUNC__ADDCOMPMARKETS,
                Arrays.<Type>asList(new DynamicArray<Address>(
                        Address.class,
                        org.web3j.abi.Utils.typeMap(cTokens, Address.class))),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _become(String unitroller) {
        final Function function = new Function(
                FUNC__BECOME,
                Arrays.<Type>asList(new Address(160, unitroller)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> _borrowGuardianPaused() {
        final Function function = new Function(FUNC__BORROWGUARDIANPAUSED,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> _dropCompMarket(String cToken) {
        final Function function = new Function(
                FUNC__DROPCOMPMARKET,
                Arrays.<Type>asList(new Address(160, cToken)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> _mintGuardianPaused() {
        final Function function = new Function(FUNC__MINTGUARDIANPAUSED,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> _setBorrowCapGuardian(String newBorrowCapGuardian) {
        final Function function = new Function(
                FUNC__SETBORROWCAPGUARDIAN,
                Arrays.<Type>asList(new Address(160, newBorrowCapGuardian)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setBorrowPaused(String cToken, Boolean state) {
        final Function function = new Function(
                FUNC__SETBORROWPAUSED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Bool(state)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setCloseFactor(BigInteger newCloseFactorMantissa) {
        final Function function = new Function(
                FUNC__SETCLOSEFACTOR,
                Arrays.<Type>asList(new Uint256(newCloseFactorMantissa)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setCollateralFactor(String cToken, BigInteger newCollateralFactorMantissa) {
        final Function function = new Function(
                FUNC__SETCOLLATERALFACTOR,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Uint256(newCollateralFactorMantissa)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setCompRate(BigInteger compRate_) {
        final Function function = new Function(
                FUNC__SETCOMPRATE,
                Arrays.<Type>asList(new Uint256(compRate_)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setLiquidationIncentive(BigInteger newLiquidationIncentiveMantissa) {
        final Function function = new Function(
                FUNC__SETLIQUIDATIONINCENTIVE,
                Arrays.<Type>asList(new Uint256(newLiquidationIncentiveMantissa)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setMarketBorrowCaps(List<String> cTokens, List<BigInteger> newBorrowCaps) {
        final Function function = new Function(
                FUNC__SETMARKETBORROWCAPS,
                Arrays.<Type>asList(new DynamicArray<Address>(
                                Address.class,
                                org.web3j.abi.Utils.typeMap(cTokens, Address.class)),
                        new DynamicArray<Uint256>(
                                Uint256.class,
                                org.web3j.abi.Utils.typeMap(newBorrowCaps, Uint256.class))),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setMaxAssets(BigInteger newMaxAssets) {
        final Function function = new Function(
                FUNC__SETMAXASSETS,
                Arrays.<Type>asList(new Uint256(newMaxAssets)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setMintPaused(String cToken, Boolean state) {
        final Function function = new Function(
                FUNC__SETMINTPAUSED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Bool(state)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setPauseGuardian(String newPauseGuardian) {
        final Function function = new Function(
                FUNC__SETPAUSEGUARDIAN,
                Arrays.<Type>asList(new Address(160, newPauseGuardian)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setPriceOracle(String newOracle) {
        final Function function = new Function(
                FUNC__SETPRICEORACLE,
                Arrays.<Type>asList(new Address(160, newOracle)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setSeizePaused(Boolean state) {
        final Function function = new Function(
                FUNC__SETSEIZEPAUSED,
                Arrays.<Type>asList(new Bool(state)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _setTransferPaused(Boolean state) {
        final Function function = new Function(
                FUNC__SETTRANSFERPAUSED,
                Arrays.<Type>asList(new Bool(state)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> _supportMarket(String cToken) {
        final Function function = new Function(
                FUNC__SUPPORTMARKET,
                Arrays.<Type>asList(new Address(160, cToken)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<String> accountAssets(String param0, BigInteger param1) {
        final Function function = new Function(FUNC_ACCOUNTASSETS,
                Arrays.<Type>asList(new Address(160, param0),
                        new Uint256(param1)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<String> admin() {
        final Function function = new Function(FUNC_admin,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<String> allMarkets(BigInteger param0) {
        final Function function = new Function(FUNC_ALLMARKETS,
                Arrays.<Type>asList(new Uint256(param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<TransactionReceipt> borrowAllowed(String cToken, String borrower, BigInteger borrowAmount) {
        final Function function = new Function(
                FUNC_BORROWALLOWED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, borrower),
                        new Uint256(borrowAmount)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<String> borrowCapGuardian() {
        final Function function = new Function(FUNC_BORROWCAPGUARDIAN,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<BigInteger> borrowCaps(String param0) {
        final Function function = new Function(FUNC_BORROWCAPS,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<Boolean> borrowGuardianPaused(String param0) {
        final Function function = new Function(FUNC_BORROWGUARDIANPAUSED,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> borrowVerify(String cToken, String borrower, BigInteger borrowAmount) {
        final Function function = new Function(
                FUNC_BORROWVERIFY,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, borrower),
                        new Uint256(borrowAmount)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> checkMembership(String account, String cToken) {
        final Function function = new Function(FUNC_CHECKMEMBERSHIP,
                Arrays.<Type>asList(new Address(160, account),
                        new Address(160, cToken)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> claimComp(String holder, List<String> cTokens) {
        final Function function = new Function(
                FUNC_claimComp,
                Arrays.<Type>asList(new Address(160, holder),
                        new DynamicArray<Address>(
                                Address.class,
                                org.web3j.abi.Utils.typeMap(cTokens, Address.class))),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> claimComp(List<String> holders, List<String> cTokens, Boolean borrowers, Boolean suppliers) {
        final Function function = new Function(
                FUNC_claimComp,
                Arrays.<Type>asList(new DynamicArray<Address>(
                                Address.class,
                                org.web3j.abi.Utils.typeMap(holders, Address.class)),
                        new DynamicArray<Address>(
                                Address.class,
                                org.web3j.abi.Utils.typeMap(cTokens, Address.class)),
                        new Bool(borrowers),
                        new Bool(suppliers)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> claimComp(String holder) {
        final Function function = new Function(
                FUNC_claimComp,
                Arrays.<Type>asList(new Address(160, holder)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<BigInteger> closeFactorMantissa() {
        final Function function = new Function(FUNC_CLOSEFACTORMANTISSA,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<BigInteger> compAccrued(String param0) {
        final Function function = new Function(FUNC_COMPACCRUED,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<Tuple2<BigInteger, BigInteger>> compBorrowState(String param0) {
        final Function function = new Function(FUNC_COMPBORROWSTATE,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint224>() {
                }, new TypeReference<Uint32>() {
                }));
        return new RemoteFunctionCall<Tuple2<BigInteger, BigInteger>>(function,
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(),
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteFunctionCall<BigInteger> compBorrowerIndex(String param0, String param1) {
        final Function function = new Function(FUNC_COMPBORROWERINDEX,
                Arrays.<Type>asList(new Address(160, param0),
                        new Address(160, param1)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<BigInteger> compClaimThreshold() {
        final Function function = new Function(FUNC_COMPCLAIMTHRESHOLD,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<BigInteger> compInitialIndex() {
        final Function function = new Function(FUNC_COMPINITIALINDEX,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint224>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<BigInteger> compRate() {
        final Function function = new Function(FUNC_COMPRATE,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<BigInteger> compSpeeds(String param0) {
        final Function function = new Function(FUNC_COMPSPEEDS,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<BigInteger> compSupplierIndex(String param0, String param1) {
        final Function function = new Function(FUNC_COMPSUPPLIERINDEX,
                Arrays.<Type>asList(new Address(160, param0),
                        new Address(160, param1)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<Tuple2<BigInteger, BigInteger>> compSupplyState(String param0) {
        final Function function = new Function(FUNC_COMPSUPPLYSTATE,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint224>() {
                }, new TypeReference<Uint32>() {
                }));
        return new RemoteFunctionCall<Tuple2<BigInteger, BigInteger>>(function,
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(),
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteFunctionCall<String> comptrollerImplementation() {
        final Function function = new Function(FUNC_comptrollerImplementation,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<TransactionReceipt> enterMarkets(List<String> cTokens) {
        final Function function = new Function(
                FUNC_ENTERMARKETS,
                Arrays.<Type>asList(new DynamicArray<Address>(
                        Address.class,
                        org.web3j.abi.Utils.typeMap(cTokens, Address.class))),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> exitMarket(String cTokenAddress) {
        final Function function = new Function(
                FUNC_EXITMARKET,
                Arrays.<Type>asList(new Address(160, cTokenAddress)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Tuple3<BigInteger, BigInteger, BigInteger>> getAccountLiquidity(String account) {
        final Function function = new Function(FUNC_GETACCOUNTLIQUIDITY,
                Arrays.<Type>asList(new Address(160, account)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }, new TypeReference<Uint256>() {
                }, new TypeReference<Uint256>() {
                }));
        return new RemoteFunctionCall<Tuple3<BigInteger, BigInteger, BigInteger>>(function,
                new Callable<Tuple3<BigInteger, BigInteger, BigInteger>>() {
                    @Override
                    public Tuple3<BigInteger, BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple3<BigInteger, BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(),
                                (BigInteger) results.get(1).getValue(),
                                (BigInteger) results.get(2).getValue());
                    }
                });
    }

    public RemoteFunctionCall<List> getAllMarkets() {
        final Function function = new Function(FUNC_GETALLMARKETS,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<DynamicArray<Address>>() {
                }));
        return new RemoteFunctionCall<List>(function,
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteFunctionCall<List> getAssetsIn(String account) {
        final Function function = new Function(FUNC_GETASSETSIN,
                Arrays.<Type>asList(new Address(160, account)),
                Arrays.<TypeReference<?>>asList(new TypeReference<DynamicArray<Address>>() {
                }));
        return new RemoteFunctionCall<List>(function,
                new Callable<List>() {
                    @Override
                    @SuppressWarnings("unchecked")
                    public List call() throws Exception {
                        List<Type> result = (List<Type>) executeCallSingleValueReturn(function, List.class);
                        return convertToNative(result);
                    }
                });
    }

    public RemoteFunctionCall<BigInteger> getBlockNumber() {
        final Function function = new Function(FUNC_GETBLOCKNUMBER,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<String> getCompAddress() {
        final Function function = new Function(FUNC_GETCOMPADDRESS,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<Tuple3<BigInteger, BigInteger, BigInteger>> getHypotheticalAccountLiquidity(String account, String cTokenModify, BigInteger redeemTokens, BigInteger borrowAmount) {
        final Function function = new Function(FUNC_GETHYPOTHETICALACCOUNTLIQUIDITY,
                Arrays.<Type>asList(new Address(160, account),
                        new Address(160, cTokenModify),
                        new Uint256(redeemTokens),
                        new Uint256(borrowAmount)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }, new TypeReference<Uint256>() {
                }, new TypeReference<Uint256>() {
                }));
        return new RemoteFunctionCall<Tuple3<BigInteger, BigInteger, BigInteger>>(function,
                new Callable<Tuple3<BigInteger, BigInteger, BigInteger>>() {
                    @Override
                    public Tuple3<BigInteger, BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple3<BigInteger, BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(),
                                (BigInteger) results.get(1).getValue(),
                                (BigInteger) results.get(2).getValue());
                    }
                });
    }

    public RemoteFunctionCall<Boolean> isComptroller() {
        final Function function = new Function(FUNC_ISCOMPTROLLER,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> liquidateBorrowAllowed(String cTokenBorrowed, String cTokenCollateral, String liquidator, String borrower, BigInteger repayAmount) {
        final Function function = new Function(
                FUNC_LIQUIDATEBORROWALLOWED,
                Arrays.<Type>asList(new Address(160, cTokenBorrowed),
                        new Address(160, cTokenCollateral),
                        new Address(160, liquidator),
                        new Address(160, borrower),
                        new Uint256(repayAmount)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> liquidateBorrowVerify(String cTokenBorrowed, String cTokenCollateral, String liquidator, String borrower, BigInteger actualRepayAmount, BigInteger seizeTokens) {
        final Function function = new Function(
                FUNC_LIQUIDATEBORROWVERIFY,
                Arrays.<Type>asList(new Address(160, cTokenBorrowed),
                        new Address(160, cTokenCollateral),
                        new Address(160, liquidator),
                        new Address(160, borrower),
                        new Uint256(actualRepayAmount),
                        new Uint256(seizeTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Tuple2<BigInteger, BigInteger>> liquidateCalculateSeizeTokens(String cTokenBorrowed, String cTokenCollateral, BigInteger actualRepayAmount) {
        final Function function = new Function(FUNC_LIQUIDATECALCULATESEIZETOKENS,
                Arrays.<Type>asList(new Address(160, cTokenBorrowed),
                        new Address(160, cTokenCollateral),
                        new Uint256(actualRepayAmount)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }, new TypeReference<Uint256>() {
                }));
        return new RemoteFunctionCall<Tuple2<BigInteger, BigInteger>>(function,
                new Callable<Tuple2<BigInteger, BigInteger>>() {
                    @Override
                    public Tuple2<BigInteger, BigInteger> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple2<BigInteger, BigInteger>(
                                (BigInteger) results.get(0).getValue(),
                                (BigInteger) results.get(1).getValue());
                    }
                });
    }

    public RemoteFunctionCall<BigInteger> liquidationIncentiveMantissa() {
        final Function function = new Function(FUNC_LIQUIDATIONINCENTIVEMANTISSA,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<Tuple3<Boolean, BigInteger, Boolean>> markets(String param0) {
        final Function function = new Function(FUNC_MARKETS,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }, new TypeReference<Uint256>() {
                }, new TypeReference<Bool>() {
                }));
        return new RemoteFunctionCall<Tuple3<Boolean, BigInteger, Boolean>>(function,
                new Callable<Tuple3<Boolean, BigInteger, Boolean>>() {
                    @Override
                    public Tuple3<Boolean, BigInteger, Boolean> call() throws Exception {
                        List<Type> results = executeCallMultipleValueReturn(function);
                        return new Tuple3<Boolean, BigInteger, Boolean>(
                                (Boolean) results.get(0).getValue(),
                                (BigInteger) results.get(1).getValue(),
                                (Boolean) results.get(2).getValue());
                    }
                });
    }

    public RemoteFunctionCall<BigInteger> maxAssets() {
        final Function function = new Function(FUNC_MAXASSETS,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {
                }));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<TransactionReceipt> mintAllowed(String cToken, String minter, BigInteger mintAmount) {
        final Function function = new Function(
                FUNC_MINTALLOWED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, minter),
                        new Uint256(mintAmount)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> mintGuardianPaused(String param0) {
        final Function function = new Function(FUNC_MINTGUARDIANPAUSED,
                Arrays.<Type>asList(new Address(160, param0)),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> mintVerify(String cToken, String minter, BigInteger actualMintAmount, BigInteger mintTokens) {
        final Function function = new Function(
                FUNC_MINTVERIFY,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, minter),
                        new Uint256(actualMintAmount),
                        new Uint256(mintTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<String> oracle() {
        final Function function = new Function(FUNC_ORACLE,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<String> pauseGuardian() {
        final Function function = new Function(FUNC_PAUSEGUARDIAN,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<String> pendingAdmin() {
        final Function function = new Function(FUNC_pendingAdmin,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<String> pendingComptrollerImplementation() {
        final Function function = new Function(FUNC_pendingComptrollerImplementation,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {
                }));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<TransactionReceipt> redeemAllowed(String cToken, String redeemer, BigInteger redeemTokens) {
        final Function function = new Function(
                FUNC_REDEEMALLOWED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, redeemer),
                        new Uint256(redeemTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> redeemVerify(String cToken, String redeemer, BigInteger redeemAmount, BigInteger redeemTokens) {
        final Function function = new Function(
                FUNC_REDEEMVERIFY,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, redeemer),
                        new Uint256(redeemAmount),
                        new Uint256(redeemTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> refreshCompSpeeds() {
        final Function function = new Function(
                FUNC_REFRESHCOMPSPEEDS,
                Arrays.<Type>asList(),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> repayBorrowAllowed(String cToken, String payer, String borrower, BigInteger repayAmount) {
        final Function function = new Function(
                FUNC_REPAYBORROWALLOWED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, payer),
                        new Address(160, borrower),
                        new Uint256(repayAmount)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> repayBorrowVerify(String cToken, String payer, String borrower, BigInteger actualRepayAmount, BigInteger borrowerIndex) {
        final Function function = new Function(
                FUNC_REPAYBORROWVERIFY,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, payer),
                        new Address(160, borrower),
                        new Uint256(actualRepayAmount),
                        new Uint256(borrowerIndex)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> seizeAllowed(String cTokenCollateral, String cTokenBorrowed, String liquidator, String borrower, BigInteger seizeTokens) {
        final Function function = new Function(
                FUNC_SEIZEALLOWED,
                Arrays.<Type>asList(new Address(160, cTokenCollateral),
                        new Address(160, cTokenBorrowed),
                        new Address(160, liquidator),
                        new Address(160, borrower),
                        new Uint256(seizeTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> seizeGuardianPaused() {
        final Function function = new Function(FUNC_SEIZEGUARDIANPAUSED,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> seizeVerify(String cTokenCollateral, String cTokenBorrowed, String liquidator, String borrower, BigInteger seizeTokens) {
        final Function function = new Function(
                FUNC_SEIZEVERIFY,
                Arrays.<Type>asList(new Address(160, cTokenCollateral),
                        new Address(160, cTokenBorrowed),
                        new Address(160, liquidator),
                        new Address(160, borrower),
                        new Uint256(seizeTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> transferAllowed(String cToken, String src, String dst, BigInteger transferTokens) {
        final Function function = new Function(
                FUNC_TRANSFERALLOWED,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, src),
                        new Address(160, dst),
                        new Uint256(transferTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> transferGuardianPaused() {
        final Function function = new Function(FUNC_TRANSFERGUARDIANPAUSED,
                Arrays.<Type>asList(),
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {
                }));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> transferVerify(String cToken, String src, String dst, BigInteger transferTokens) {
        final Function function = new Function(
                FUNC_TRANSFERVERIFY,
                Arrays.<Type>asList(new Address(160, cToken),
                        new Address(160, src),
                        new Address(160, dst),
                        new Uint256(transferTokens)),
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    @Deprecated
    public static ComptrollerContract load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new ComptrollerContract(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static ComptrollerContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new ComptrollerContract(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static ComptrollerContract load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new ComptrollerContract(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ComptrollerContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new ComptrollerContract(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static class NewPendingImplementationEventResponse extends BaseEventResponse {
        public String oldPendingImplementation;

        public String newPendingImplementation;
    }

    public static class NewImplementationEventResponse extends BaseEventResponse {
        public String oldImplementation;

        public String newImplementation;
    }

    public static class NewPendingAdminEventResponse extends BaseEventResponse {
        public String oldPendingAdmin;

        public String newPendingAdmin;
    }

    public static class NewAdminEventResponse extends BaseEventResponse {
        public String oldAdmin;

        public String newAdmin;
    }

    public static class FailureEventResponse extends BaseEventResponse {
        public BigInteger error;

        public BigInteger info;

        public BigInteger detail;
    }

    public static class ActionPausedEventResponse extends BaseEventResponse {
        public String cToken;

        public String action;

        public Boolean pauseState;
    }

    public static class CompSpeedUpdatedEventResponse extends BaseEventResponse {
        public String cToken;
        public BigInteger newSpeed;
    }

    public static class DistributedBorrowerCompEventResponse extends BaseEventResponse {
        public String cToken;

        public String borrower;

        public BigInteger compDelta;

        public BigInteger compBorrowIndex;
    }

    public static class DistributedSupplierCompEventResponse extends BaseEventResponse {
        public String cToken;

        public String supplier;

        public BigInteger compDelta;

        public BigInteger compSupplyIndex;
    }

    public static class MarketCompedEventResponse extends BaseEventResponse {
        public String cToken;

        public Boolean isComped;
    }

    public static class MarketEnteredEventResponse extends BaseEventResponse {
        public String cToken;

        public String account;
    }

    public static class MarketExitedEventResponse extends BaseEventResponse {
        public String cToken;

        public String account;
    }

    public static class MarketListedEventResponse extends BaseEventResponse {
        public String cToken;
    }

    public static class NewBorrowCapEventResponse extends BaseEventResponse {
        public String cToken;

        public BigInteger newBorrowCap;
    }

    public static class NewBorrowCapGuardianEventResponse extends BaseEventResponse {
        public String oldBorrowCapGuardian;

        public String newBorrowCapGuardian;
    }

    public static class NewCloseFactorEventResponse extends BaseEventResponse {
        public BigInteger oldCloseFactorMantissa;

        public BigInteger newCloseFactorMantissa;
    }

    public static class NewCollateralFactorEventResponse extends BaseEventResponse {
        public String cToken;

        public BigInteger oldCollateralFactorMantissa;

        public BigInteger newCollateralFactorMantissa;
    }

    public static class NewCompRateEventResponse extends BaseEventResponse {
        public BigInteger oldCompRate;

        public BigInteger newCompRate;
    }

    public static class NewLiquidationIncentiveEventResponse extends BaseEventResponse {
        public BigInteger oldLiquidationIncentiveMantissa;

        public BigInteger newLiquidationIncentiveMantissa;
    }

    public static class NewMaxAssetsEventResponse extends BaseEventResponse {
        public BigInteger oldMaxAssets;

        public BigInteger newMaxAssets;
    }

    public static class NewPauseGuardianEventResponse extends BaseEventResponse {
        public String oldPauseGuardian;

        public String newPauseGuardian;
    }

    public static class NewPriceOracleEventResponse extends BaseEventResponse {
        public String oldPriceOracle;

        public String newPriceOracle;
    }
}

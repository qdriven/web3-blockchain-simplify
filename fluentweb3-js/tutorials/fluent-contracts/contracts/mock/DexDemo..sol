pragma solidity >=0.8

contract DexDemo{
    struct Token{
        bytes32 ticker;
        address tokenAddress;
    }

    mapping(bytes32=>Token) public tokens;
    bytes32[] public tokenList;
    mapping(address=>mapping(bytes32=>uint)) public traderBalances;
    constructor() public {
        admin=msg.sender;
    }

    /**
    Add ticker and token
     */
    function addToken(bytes32 ticker,address tokenAddress) onlyAdmin() external{
        tokens[ticker]=Token(ticker,tokenAddress)
        tokenList.push(ticker)
    }

    function deposit(uint amount,bytes32 ticker) tokenExist(ticker) external{
        IERC20(tokens[ticker].tokenAddress).transferFrom(
            msg.sender,
            address(this),
            amount
        )
        traderBalances[msg.sender][ticker] += amount;
    }
  function withdraw(
        uint amount,
        bytes32 ticker)
        tokenExist(ticker)
        external {
        require(
            traderBalances[msg.sender][ticker] >= amount,
            'balance too low'
        ); 
        traderBalances[msg.sender][ticker] -= amount;
        IERC20(tokens[ticker].tokenAddress).transfer(msg.sender, amount);
    }
      modifier tokenExist(bytes32 ticker) {
        require(
            tokens[ticker].tokenAddress != address(0),
            'this token does not exist'
        );
        _;
    }
     modifier onlyAdmin() {
        require(msg.sender == admin, 'only admin');
        _;
    }
}
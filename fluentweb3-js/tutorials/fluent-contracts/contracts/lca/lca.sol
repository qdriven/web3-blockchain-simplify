pragma solidity >=0.8

contract CarbonFootPrint {
    struct Product{
        uint32 id;
        string name;
        string description;
        bool intermediate;
        uint32 idOrganization; //ref to Org.
        uint32 idUnit; //reference to Unit
        uint32[] productFootPrints;
        //refs to ProductFootprint
    }
    struct ProductFootprint{
        uint32 id;
        uint32 co2eq; // base
        uint16 exp; // exponent
        uint32 idProduct;
        uint32 year;
        string month;
        uint32 idMA; //ref to MonthlyActivity
    }
    //-- Structures - Entities Implementation --
    struct MonthlyActivity{
        uint32 id;
        string description;
        uint32 co2e;
        uint16 exp;
        string month;
        uint32[] productQuantities;
        //refs to input prod quantities
        uint32 output; //ref to output Prod FootPrint
        uint32 finalProductQty;
        uint32 idOrganization;
        uint32 idUnit;
        uint32 idYear;
        address idUser;
        unit32[] productionCosts;
    }
    struct MonthlyFixedCost{
        uint32 id;
        uint32 co2e;
        uint16 exp;
        string description;
        uint32 quantity;
        string month;
        uint32 idCostType;
        uint32 idOrganization;
        uint32 idYear;
    }
    mapping(uint32 => Product) public products;
    uint32 public productsCount;
    mapping(uint32 => MonthlyActivity)
    public mactivities;
    uint32 public mactivitiesCount;
    mapping(uint32 => Organization)
    public organizations;
    uint32 public organizationsCount;
    mapping(uint32 => MonthlyFixCost)
    public mfixcosts;
    uint32 public mfixcostsCount;
    mapping(uint32 => Unit) public units;
    uint32 public unitsCount;
    mapping(uint32 => ProductCost)
    public productCosts;
    uint32 public productCostsCount;
    mapping(uint32 => CostType)
    public costsTypes;
    uint32 public costsTypesCount;
    mapping(uint32 => ProductFootprint)
    public productFootPrints;
    uint32 public pfootPrintCount;
    mapping(uint32 => ProductQuantity)
    public productsQuantities;
    uint32 public productsQuantitiesCount

    event registUserEvent (address indexed _candidateAddress);
    constructor () public{
            users[msg.sender] = User(msg.sender, 0, true);
            arrayUsers.push(msg.sender);
            // -- Initalize units
            addUnit("tonne", "t", 10, 0, 1, false);
            addUnit("kilogram", "kg", 10, 3, 1, true);
            addUnit("gram", "g", 10, 6, 1, true);
            addUnit("milligram", "mg", 10, 9, 1, true);
        }

    function addUser (address _userResp,address _user, uint16 _tipo,uint32 _organization) public {
        require(users[_user].user_add == address(0),"User already registered");
        if(_tipo == 0 || _tipo == 1){
            require(users[_userResp].tipo == 0,"You need admin permissions");
        }else if(_tipo == 2){
            require(users[_userResp].tipo == 1 ,"You need organization admin permissions");
        }
        users[_user] = User(_user, _tipo, true);
        arrayUsers.push(_user);
        userOrganizations[_user].
        push(_organization);
        emit registUserEvent(_user);
    }

    function addProduct(string memory _name,string memory _description,bool _intermediate,uint32 _org,
        uint32 _unit,uint32[] memory _footPrints) public{
            bool exist = false;
        require(users[msg.sender].idOrganization == _org,"You need to belong to the organization");
        require(organizations[_org].id != uint32(0),"Organization doesn’t exist");
        for(uint32 i=1; i <= productsCount; i++){
            string memory name = products[i].name;
            if(keccak256(abi.encodePacked(name)) == keccak256(abi.encodePacked(_name))){
                exist = true;
            }
        }
        require(!exist, "Product already registered");
        productsCount++;    
        products[productsCount] =Product(productsCount, _name, _description, units[_unit].initials,
            _intermediate, _org, _unit, _footPrints);
            organizations[_org].products.push(productsCount);
    }
}
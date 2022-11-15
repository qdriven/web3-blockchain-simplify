## 1.Energy Web是什么

>Energy Web’s mission is to accelerate decarbonization of the global economy.
To do so, we deploy digital operating systems for energy grids with our global community of more than 100 energy market participants. These systems make it simple, secure, and efficient for clean energy assets to support the grid of the future.

Energy Web目前主要针对的是电网，他们认为的问题是:
1.  2030年，电力市场会更加的消费者-中心化，目前的电网中会有各种renewable energy, distributed energy resources (DERs), and electric mobility
2. 目前电网的问题，他是假设电网供给是可控的, 需求是固定的
3. 他们认为电力不会像现在这样的中心化提供
4. E-Mobility/EVs问题，EV什么时候，如何进行充电对电网有非常大的影响,
5. **EV使用的电是否是绿电还是非绿电，没有办法追踪**

https://www.energyweb.org/industries/grid-operators/

总体感觉针对的问题是：
1. 电力提供方和消费方会更C端化，目前电网不支持这些
2. 绿电和化石能源发电在进入电网里面后目前技术没有办法区分，需要追踪能源生成的情况
3. 追踪电力生成/消费情况，需要数据真实和可验证，所以采用区块链方案+去中心化身份(DID)+智能合约方案
4. 目前的能源属性的认证都是在各个不同机构，效率低，可能不能满足后面的大众化需求，需要一个公开的网络进行多方协作，所以采用区块链

采用的方案:
1. 底层可信数据存储，使用区块链
2. 绿电能源生成接入, 使用去中心化身份(DIDs and Verifiable Credentials)的方式区别不同设别，
   流转证明文件
3. 证明文件，用来证明实际产生的数据，连接认证机构+设备来验证确认数据证明，在将证明和数据通过NFT(双Token协议)进行数字化，从而进行交易

从最近的一些POC来看，看起来对于电动车充电比较看重.

[网站文档](https://energy-web-foundation.gitbook.io/)

## 2.Energy Web  目标问题/Use Case/解决方案

## 2.1 Green Proofs
Energey Web的解决方案主要是： [green-proofs](https://www.energyweb.org/solutions/green-proofs/),主要用来：
> tracking low-carbon products and their attributes in a highly scalable, transparent, and efficient manner

Energy Web 区块链使用的技术主要包括了：
1. DID: decentralized identifiers, verifiable credentials
2. business logic embedded in smart contracts

主要的目标: 追踪碳排放，确认每一个单位或者产品事实是GREEN
> Green Proofs makes it easy to establish a transparent audit trail proving that a given unit of energy or a given product (e.g., electricity, hydrogen, carbon offsets, aviation fuel, green steel) is in fact “green.”

计划： Green Proofs发布在2022年Q3， 主要USE CASE:
> These applications are configured around three primary use cases: 24/7 matched renewable energy, decarbonized bitcoin, and verified renewable electric vehicle charging.

## 2.2 Facilitating Clean Energy Purchases

清洁能源购买的目前的一个挑战:
1. 通过清洁能源生产的电力能源，是使用中(电网中)没有办法区别出来,和石化的发电没有办法区别
2. EAC(Energy Attribute Certification), 用来进行对能源进行追踪和打标签，甚至交易
> EACs have different names in different markets, including guarantees of origin (GOs), renewable energy certificates (RECs), and solar renewable energy certificates (S-RECs). These are conceptually similar.


Energy Web 提供解决方案解决以下几个问题:
1. Application 1: **Decentralizing and Automating** [**Energy Attribute Certificate**](/energy-web/glossary-of-terms#energy-attribute-certificates-eacs) **Issuance**
2. Application 2: Facilitating Transparent and Competitive [**Energy Attribute Certificate**](/energy-web/glossary-of-terms#energy-attribute-certificates-eacs) **Trading**
3. Application 3: Provide end-to-end infrastructure for a **renewable energy procurement system**

这块主要的想解决的问题是:
1.  EAC市场比较封闭和碎片化,有很多的专业机构，目前缺少合作
2. EAC进入市场成本很高，目前有很多的管理成本存在，认证信息主要存放在机构数据库中
3. 现有机构的系统比较老旧，自动化程度低，EAC认证时间长成本高
4. 在EAC交易层面，目前没有统一市场，都是通过代理交易，需要一个公开市场，让价格透明
![[rec.png]]

## 2.3  Scaling Access to Grid Flexibility

目前存在的问题：
1. [Distributed Energy Resources (DERs)](/energy-web/glossary-of-terms#distributed-energy-der-resources)  会增多，现有电网不能支持，电网接入设备会增多和传统电网架构不匹配，DER是相对小型设备，可能同时消费和生产电力(比如: rooftop solar panels, home batteries, and electric vehicles)
 2. 现有电网数字化/可编程性不够，不能动态调整价格，容量，地址等，没有很好的load balancing的作用

Energy Web 提供解决方案解决以下几个问题:
1. Application 1: [Prosumer](/energy-web/glossary-of-terms#prosumer) Coordination
2. Application 2: Emergency Demand Flexibility
3. Application 3: **Application and IoT Management**
4. Applied Use Case: **Battery Recycling in Belgium (EasyBat)**

总结以上的解决方案，重点在于:
1.  可以让DER 设备快速接入电网，通过协议/IOT设备等，让电网了解设备的真实情况
2.  为了真实起到低碳的作用，这些设别，比如电池/电池回收等，需要给电池有一个数字身份(digital passport), 目前数字身份主要采用了:
	- [DIDs and Verifiable Credentials](https://github.com/energywebfoundation/ewf-gitbook/blob/develop/foundational-concepts/self-sovereign-identity/README.md#components-of-ssi).
	- 设备生命周期数据可以尝试存放到 IPFS这种分布式存储上，用作数据可信证明
![[did.png]]

##  3. Energy Web Tech Stack

Energy Web的技术栈结构 ：

![[ew-dos-techstack.png]]

1. Trust Layer: Blockchain/Ledger, 类似以太坊公链，有智能合约，数据不可更改
	1. 区块链治理机制: PoA(, [**Proof-of-Authority**](https://openethereum.github.io/Proof-of-Authority-Chains))共识算法，也就是区块链验证节点(validator)需要认证
	2. 区块链验证节点需要是注册的企业，具有技术能力，并且是EWF(energy-web-foundation) 的会员
2. Utility 主要用途： Identity and Access Management (IAM)，数字身份
	1. DID存储，身份认证，设备的数字身份，同时包含
	2. 交易信息保存
3. Applicaitons/SDK: 主要针对不同的应用场景的组件，构建在Blockchain和Utility层上，主要聚焦场景:
	-  Scaling access to grid flexibility
	- Facilitating clean energy purchases


### Application/SDK 详细

#### Applicaition/SDK总体模块

1. 数字身份模块: Switchboard
2. 溯源模块: Origin(Core,24/7,Zero,Local Marketplace)

![[application-sdk.png]]

####  Clean Energy Traceability Applications and SDKs - Origin
building applications that support issuance, trading, tracking, cancellations / retirements, and claims reporting ***[Energy Attribute Certificates (EACs)](https://energy-web-foundation.gitbook.io/energy-web/glossary-of-terms#energy-attribute-certificates-eacs)***

#### Origin Core SDK

> Origin allows low-carbon energy suppliers to sell their output in the form of [energy attribute certificates (EACs)](/energy-web/glossary-of-terms#energy-attribute-certificates-eacs), and allows consumers to purchase these certificates in a peer-to-peer marketplace. The data and lifecycle of the EAC is tracked on the Energy Web Chain, providing transparency to all stakeholders involved in the process.

Origin的作用是把低碳能源以一种EAC的方式进行买卖. EAC可以在一个点对点，或者相对开放的市场进行交易, EAC的生命周期情况通过Energy Web Chain进行监控. 

下面是Origin可以来做的: 
1. EAC 的 Issue： 数字化/代币化
2. EAC 交易市场： trade, claim, report ，同时能够和现有的机构系统进行直连
3. 能源设备可以自动通过Energy Web注册进入EAC市场
4. 支持不同的认证标准:  REC, Guarantees of Origin, I-REC

Origin的模块:
1. [**Registry (device) modules**](https://github.com/energywebfoundation/origin/tree/master/packages/devices): 设备注册，包括设备拥有人的信息
2. [**Issuer (traceability) modules**](https://github.com/energywebfoundation/origin/tree/master/packages/traceability)：跟踪设备EAC的使用情况，同时用收集的数据和认证机构系统保持同步
3. [**Exchange (trade) modules**](https://github.com/energywebfoundation/origin/tree/master/packages/trade)： 交易EAC系统

#### Origin Issue模块
**Energy Web Origin's Issuer module addresses this pain point by enabling an open and automated issuance process that is compatible with any certification standard**.
EAC的Issue发行这块，主要使用了: [ERC-1888 Transferable Certificate Claim](https://github.com/ethereum/EIPs/issues/1888) (which is an extension of the [ERC-1155 Multi token Standard](https://eips.ethereum.org/EIPS/eip-1155)) 标准，实际情况类似于如下流程:

1.  EAC 设备方发起申请(Origin's Device Registration Module)
2. Issuer module 获取数据验证，数据可以通过以下方式获得:
	1.  发起方自动输入
	2. 直接通过API或者IoT设别活的
3.   Issuer module 将获得的设备数据发送到认证机构进行认证审查
4.  Issuer module模块通过认证机构数据库获取认证是否已经通过的信息
5. 如果通过了，那么就发布一个ERC-1888的token， 这个token实际也是一种NFTtoken，具有唯一性，同时具有数量
	> The ERC-1155 standard supports fungible and non-fungible components. The device’s kilowatt hours generated are stored as fungible tokens that represent certified green energy. Each hour or combination of hours can be split and sold to different customers. The device information, total energy volume and timeframe of generation are stored as an NFT asset, which cannot be split.
 ![[erc-1888.png]]
 这块和Toucan 协议不同的地方是，直接用ERC-1888这个multiple token标准解决代币化的过程
 - Non-Fungible部分代表了 设备或者项目
 - Fungible 部分代表可以了可以交易的数量

认证信息的数据结构 :

```
struct Certificate {
    uint256 topic;
    address issuer; // msg.sender
    bytes validityData;
    bytes issuenceData;
}
```
Topic数据可以为:
```
Possible certificate topics:

-   `01`: Utilities
    -   `01`: Electricity
        -   `01`: Origin ("Green certificates")
        -   `02`: Origin + Electricity
        -   `03`: Flexibility
        -   `04`: Flexibility + Electricity
        -   `05`: Efficiency ("White certificates")
    -   `02`: Water
    -   `03`: Gas
    -   `04`: Heat
```


### 3.2 Identity Management Applications

Switchboard: DID/IAM 身份权限管理，
-   Create self-sovereign [D](/energy-web/identity-at-ew/self-sovereign-identity-introduction#decentralized-identities-dids)[ecentralized Identities](/energy-web/identity-at-ew/self-sovereign-identity-introduction#decentralized-identities-dids) (DID) using a connection with a crypto wallet like MetaMask. This DID allows assets to participate in the EW-DOS system.
-   Define structures for organizations, applications and role-based permissions to participate in them
-   Request and issue [verified credentials](/energy-web/identity-at-ew/self-sovereign-identity-introduction#verifiable-credentials-vcs) that are required to take on roles within an organization or an application
关于DID/Verified Credential的说明: 个人理解大致就是
1.  DID：  [D](/energy-web/identity-at-ew/self-sovereign-identity-introduction#decentralized-identities-dids)[ecentralized Identities](/energy-web/identity-at-ew/self-sovereign-identity-introduction#decentralized-identities-dids) Documentation， 就是通过一个id把确定对应信息存放的位置，id体系就是区块链的public/private key体系
2. 各种可以验证的信息，可以在认证机构使用私钥签名的方式保存到和这个id相关联的位置，
3. 通过id可以获取所有的相关认证信息，同时出示给其他人进行证明
![[did-vc.png]]

> An example of DIDs and VCs in an asset lifecycle[](#an-example-of-dids-and-vcs-in-an-asset-lifecycle)
To understand how and why an asset can take on a DID and acquire and use VCs using EW-DOS, let’s use the lifecycle of a battery as an example.
   1. When the battery is manufactured, the manufacturer assigns a DID to the battery, and issues claims on the battery’s physical makeup (date of manufacture, the model, the serial number, the battery capacity.)
   2. The battery is sold. The installer of the battery adds new claims on the battery’s purchaser and its new location.
   3. As the battery performs, the new owner adds claims on the battery’s charge and discharge rate.
   4. 4.The battery reaches the end of its lifecycle. The final owner of the battery adds the date that it is retired, and its final charge and discharge rate.


### POC  例子

- [Volkswagen大众 Green EV Charging](https://lab.energyweb.org/greenevcharging/)
 通过连接EV充电设别，获取充电数据，可以选择时太阳能/风力/或者其他，充电完成之后生成NFT获得低碳数据
 - [Vodafone’s global connectivity platform, and Mastercard’s payment system](https://lab.energyweb.org/renewable-ev-charging/)
 >  What is Green Charging?
The green charging solution enables anyone to acquire proof that energy purchased for charging an electric vehicle comes from renewables. Using a single application with an integrated mobile wallet, users can also pick the best priced electricity and pay for it automatically.

![[green-charing.png]]

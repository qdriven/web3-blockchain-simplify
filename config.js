const userConfig = {
  title: "Web3-Blockchain In Action",
  description: "This is my awesome blog built with Flowershow",
  author: "Simon",
  // logo image
  authorLogo: "/images/logo.svg",
  // url to author website
  authorUrl: "https://qdriven.github.io/web3-blockchain-simplify",
  // links to the pages you want to link to in the navbar
  navLinks: [{ href: "/about", name: "About" }],
  // any folders/files that you want to exclude from being published; all other files in your content folder will be published
  contentExclude: ["docs/testpage.md"],
  // publish only these folders/files in your content folder (you can combine contentInclude and contentExclude)
  contentInclude: ["docs"],
  navLinks: [
    { href: "/about", name: "About" },
    //dropdown menu should not have an href and should contain 'subItems' array
    { name: "Blockchain Intro", subItems: [{ href: "/about", name: "About" }] },
    { name: "Web3 Intro", subItems: [{ href: "/about", name: "About" }] },
  ],
  navBarTitle: {
    text: "Your custom title here",
    logo: "/assets/your-logo.svg",
  },
  repoRoot: "https://github.com/flowershow/flowershow", //example
  repoEditPath: "/edit/main/site/content/", //example
  editLinkShow: false,

};

export default userConfig;

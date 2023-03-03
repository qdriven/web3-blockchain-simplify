import { expect } from "chai";

export function shouldBehaveLikeHello(): void {
  it("should return the new hello once it's changed", async function () {
    expect(await this.hello.connect(this.signers.admin).get()).to.equal(
      "Hello, world!"
    );

    await this.hello.setWorld("Bonjour, le monde!");
    expect(await this.hello.connect(this.signers.admin).get()).to.equal(
      "Bonjour, le monde!"
    );
  });
}

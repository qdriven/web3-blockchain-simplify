# Project Setup Tools

## prettier

```shell
 pnpm add prettier -d -w
 echo {}> .prettierrc.json
```

## eslint

```shell
"eslint": "^8.16.0",
"eslint-define-config": "^1.5.0",
"eslint-plugin-import": "^2.26.0",
"eslint-plugin-node": "^11.1.0",
```

## simple-git-hook
- yarn dlx simple-git-hooks
```shell
{
  "simple-git-hooks": {
    "pre-commit": "npx lint-staged",
    "pre-push": "cd ../../ && npm run format",

    // All unused hooks will be removed automatically by default
    // but you can use the `preserveUnused` option like following to prevent this behavior

    // if you'd prefer preserve all unused hooks
    "preserveUnused": true,

    // if you'd prefer preserve specific unused hooks
    "preserveUnused": ["commit-msg"]
  }
}
```

# Grammar of PDI

[![CircleCI](https://circleci.com/gh/Soontao/grammar-pdi.svg?style=shield)](https://circleci.com/gh/Soontao/grammar-pdi)
[![codecov](https://codecov.io/gh/Soontao/grammar-pdi/branch/master/graph/badge.svg)](https://codecov.io/gh/Soontao/grammar-pdi)

Provide the AST parser for BODL (Business Object Definition Language) and ABSL (ABAP Script Language)

**VSCode Plugin have some issues, sometimes developers need to delete the `.antlr` and re-generate the token/interp files**

Antlr Version: 4.7.2

## [BODL Parser](./bodl)

![](https://res.cloudinary.com/digf90pwi/image/upload/v1571731756/2019-10-22_16-08-35_rsokqx.png)

### Build

```bash
npm run build
```

sometimes we need manually edit generated lexer/parser code.
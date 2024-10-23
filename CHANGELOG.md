# Changelog

## [0.1.0-alpha](https://github.com/microserv-io/oauth2-token-vault/compare/v0.0.1-alpha...v0.1.0-alpha) (2024-10-23)


### Features

* added encryptor for encrypting sensitive fields in domain ([#2](https://github.com/microserv-io/oauth2-token-vault/issues/2)) ([c476d96](https://github.com/microserv-io/oauth2-token-vault/commit/c476d965b79ccb73b76d77978c25a724f72eb543))
* added husky for go fmt and conventional commits ([f2f4b76](https://github.com/microserv-io/oauth2-token-vault/commit/f2f4b762857a1e046510b64b807f6b1406d19aa4))
* added logging framework ([#10](https://github.com/microserv-io/oauth2-token-vault/issues/10)) ([debfb32](https://github.com/microserv-io/oauth2-token-vault/commit/debfb32a61722f8f64499c219e09721a0a230b59))
* added oauth app service ([#15](https://github.com/microserv-io/oauth2-token-vault/issues/15)) ([bd5d48c](https://github.com/microserv-io/oauth2-token-vault/commit/bd5d48cfc39bf2c2ddd2a4b483de8d39b18ed324))
* added oauth2 token source implementation ([#6](https://github.com/microserv-io/oauth2-token-vault/issues/6)) ([8dee1cc](https://github.com/microserv-io/oauth2-token-vault/commit/8dee1cc39d98ab426ccfcac4ddd38f524976b3b4))
* added release-please for release management ([c787c53](https://github.com/microserv-io/oauth2-token-vault/commit/c787c53e5030fa163e22463a5fa13aa56aa96fe0))
* added validation in config ([#8](https://github.com/microserv-io/oauth2-token-vault/issues/8)) ([7df2698](https://github.com/microserv-io/oauth2-token-vault/commit/7df2698a7d0bd31e15fce8d7c453f954e6c0d83a))
* exchange authorization code via provider service calls ([#21](https://github.com/microserv-io/oauth2-token-vault/issues/21)) ([c8c8781](https://github.com/microserv-io/oauth2-token-vault/commit/c8c87812f4b1539e2f7425c70bceb28165d40f0f))
* load database from config ([#9](https://github.com/microserv-io/oauth2-token-vault/issues/9)) ([61c93bd](https://github.com/microserv-io/oauth2-token-vault/commit/61c93bd3d886ea932202c3a51161fb9243c0a5ba))
* manage providers through gRPC ([#16](https://github.com/microserv-io/oauth2-token-vault/issues/16)) ([8e1f7de](https://github.com/microserv-io/oauth2-token-vault/commit/8e1f7de2d2f5bf78a90debdccec5a78b844cb3c3))
* oauth service implementation and connection to grpc endpoints ([#17](https://github.com/microserv-io/oauth2-token-vault/issues/17)) ([75657f5](https://github.com/microserv-io/oauth2-token-vault/commit/75657f50e045146e271060502cf9e3ccb53d001b))
* sync provider from service instead of usecase ([#27](https://github.com/microserv-io/oauth2-token-vault/issues/27)) ([721245a](https://github.com/microserv-io/oauth2-token-vault/commit/721245af64d9f22a6b05956eb266bdc0b4c2618f))


### Bug Fixes

* use pg.StringArray instead of []string as type for Gorm models ([#12](https://github.com/microserv-io/oauth2-token-vault/issues/12)) ([479b5bf](https://github.com/microserv-io/oauth2-token-vault/commit/479b5bf097e5f37a3f337720d465f80548c227f1))

# Changelog

## 0.6.0 (2026-01-18)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/ServalHQ/serval-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** manual updates ([e0e714d](https://github.com/ServalHQ/serval-go/commit/e0e714de4048dba6db4634b07be9d2a935def9bd))
* **api:** manual updates ([3c40cbd](https://github.com/ServalHQ/serval-go/commit/3c40cbd5275fba03189dfc5132f5a4c34cd347c4))
* **api:** manual updates ([0e979c5](https://github.com/ServalHQ/serval-go/commit/0e979c56f92341f6ca1d09e56feaa971cddbbe67))
* **encoder:** support bracket encoding form-data object members ([1577150](https://github.com/ServalHQ/serval-go/commit/1577150c2fa84c50dc8ce4c89abbf9927ccc95b1))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([52269d1](https://github.com/ServalHQ/serval-go/commit/52269d1697568d4df834395c34f6d723e5c34b1b))
* **mcp:** correct code tool API endpoint ([5721590](https://github.com/ServalHQ/serval-go/commit/572159066aaae06a031bf1a5098ef9db774c0c1d))
* rename param to avoid collision ([15db281](https://github.com/ServalHQ/serval-go/commit/15db281f6c28882c158264b4c34d34f1fca87a02))
* skip usage tests that don't work with Prism ([39e34a6](https://github.com/ServalHQ/serval-go/commit/39e34a67d862cd448d504574fde897a8ade3c74c))


### Chores

* add float64 to valid types for RegisterFieldValidator ([90e8b7c](https://github.com/ServalHQ/serval-go/commit/90e8b7c332f55305268c6a7dd6ba2376e0a6006c))
* bump gjson version ([c92fa8b](https://github.com/ServalHQ/serval-go/commit/c92fa8bac4f3d8dd5ec2bda1733af44a1e092a26))
* elide duplicate aliases ([2acf02d](https://github.com/ServalHQ/serval-go/commit/2acf02d3e59cacadfb0c694e3687bd0d6ef66b45))
* **internal:** codegen related update ([019173b](https://github.com/ServalHQ/serval-go/commit/019173b14a55c81b612652fef97128d1192605fd))
* **internal:** codegen related update ([ebfcd43](https://github.com/ServalHQ/serval-go/commit/ebfcd43c07ab019f352d82b4200fdde0e8d23bf8))
* **internal:** update `actions/checkout` version ([7a4f089](https://github.com/ServalHQ/serval-go/commit/7a4f089707dde8ef596a632dd85f1abd5f6d0e84))

## 0.5.0 (2025-11-10)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/ServalHQ/serval-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** manual updates ([4e73c79](https://github.com/ServalHQ/serval-go/commit/4e73c79a78f32e8f0f0853bc28233f6046f4a912))
* **api:** manual updates ([25bb480](https://github.com/ServalHQ/serval-go/commit/25bb4800c125534a054b269c65e1658106806aac))


### Chores

* **internal:** grammar fix (it's -&gt; its) ([e191ec8](https://github.com/ServalHQ/serval-go/commit/e191ec87af7f9fb84812def38f6cfe36bf8f998e))

## 0.4.0 (2025-10-16)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/ServalHQ/serval-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** manual updates ([e3298b7](https://github.com/ServalHQ/serval-go/commit/e3298b787c21673ada8085ec618b02c7b1dc754c))

## 0.3.0 (2025-09-30)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/ServalHQ/serval-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** manual updates ([267e420](https://github.com/ServalHQ/serval-go/commit/267e420ba597c966d4a72dd921cfe75985875713))
* **api:** manual updates ([5438c65](https://github.com/ServalHQ/serval-go/commit/5438c65ea97a0186c9cc6688506f8ce8df5f0c48))
* **api:** manual updates ([b5d4d1e](https://github.com/ServalHQ/serval-go/commit/b5d4d1ef72f01d993c5cf68a9a4f8aaf0d5ed04a))
* **api:** manual updates ([ab2f54b](https://github.com/ServalHQ/serval-go/commit/ab2f54be6b2b4218e53ec582265e583caa64031a))
* **api:** manual updates ([aa0e1d9](https://github.com/ServalHQ/serval-go/commit/aa0e1d96ab83db07d292f41f390db0954037f2de))

## 0.2.0 (2025-09-30)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/ServalHQ/serval-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([a086c26](https://github.com/ServalHQ/serval-go/commit/a086c2620c6957e427c7ee8aa2104fa04b66064b))
* **api:** manual updates ([4794932](https://github.com/ServalHQ/serval-go/commit/479493259218384888bb409920458c6c4f61e29d))
* **api:** update via SDK Studio ([cc250ed](https://github.com/ServalHQ/serval-go/commit/cc250ede97e670f39585bbd43839f1c92cfc2991))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([0b7a844](https://github.com/ServalHQ/serval-go/commit/0b7a8449d45309a4f57f1221acb1c69578867dac))
* **internal:** unmarshal correctly when there are multiple discriminators ([08e5ff7](https://github.com/ServalHQ/serval-go/commit/08e5ff739dfe1bc6f28edc36a0a134e12f994fd9))
* use slices.Concat instead of sometimes modifying r.Options ([93ff4c2](https://github.com/ServalHQ/serval-go/commit/93ff4c2e5197c9d25bd6122e64f3c60a119b6b0a))


### Chores

* bump minimum go version to 1.22 ([3172e67](https://github.com/ServalHQ/serval-go/commit/3172e6797947c9e77a00a4abf86ceb37250dc959))
* do not install brew dependencies in ./scripts/bootstrap by default ([030f836](https://github.com/ServalHQ/serval-go/commit/030f83693b84a4276196a531692e954428231e32))
* **internal:** codegen related update ([29a1f4c](https://github.com/ServalHQ/serval-go/commit/29a1f4c8370c63d6badfa1348add7c95e852a3c4))
* update more docs for 1.22 ([5fa4922](https://github.com/ServalHQ/serval-go/commit/5fa4922abb492a6a0c5fdd7450e798384aaaac19))

## 0.1.0 (2025-09-02)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/ServalHQ/serval-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** update via SDK Studio ([97deaf3](https://github.com/ServalHQ/serval-go/commit/97deaf3e9aeda202528fc65d2d37ef38414cc371))
* **api:** update via SDK Studio ([c11d8e2](https://github.com/ServalHQ/serval-go/commit/c11d8e2d6cb64ce898a5851b7c69986c644566dd))
* **api:** update via SDK Studio ([d4146f2](https://github.com/ServalHQ/serval-go/commit/d4146f252166edb55d142cb305d525eb401b5aa7))
* **api:** update via SDK Studio ([cd66584](https://github.com/ServalHQ/serval-go/commit/cd665845296d53d20af04bde79061fb73f2d05c9))
* **api:** update via SDK Studio ([c1fb8ad](https://github.com/ServalHQ/serval-go/commit/c1fb8ad2cd7a5185543455aa362e852b5290297d))
* **api:** update via SDK Studio ([4fccf36](https://github.com/ServalHQ/serval-go/commit/4fccf365ae1c35144901fd7deeaee6e53766b09f))
* **api:** update via SDK Studio ([067ce48](https://github.com/ServalHQ/serval-go/commit/067ce48b1988355c4b1346f733159b124e67f798))
* **api:** update via SDK Studio ([e271cf2](https://github.com/ServalHQ/serval-go/commit/e271cf2ac4c9552b2abf61f08578179ee0bf9955))
* **api:** update via SDK Studio ([4655408](https://github.com/ServalHQ/serval-go/commit/4655408d070ea338fcb63a61a2502aaa78244738))
* **api:** update via SDK Studio ([92a3cad](https://github.com/ServalHQ/serval-go/commit/92a3cadae9c5a6e3c083ba0dd8b994cdad4f230c))
* **api:** update via SDK Studio ([257687a](https://github.com/ServalHQ/serval-go/commit/257687a1613f476ae1e62a6869d54b1909ef7d60))
* **api:** update via SDK Studio ([dc76027](https://github.com/ServalHQ/serval-go/commit/dc76027e7484506e24e30888e6e14d98816afe58))
* **api:** update via SDK Studio ([e20f85c](https://github.com/ServalHQ/serval-go/commit/e20f85cfaa312e7eb590bb734bef99b92e03ad37))
* **api:** update via SDK Studio ([451b905](https://github.com/ServalHQ/serval-go/commit/451b905c33a12bcb23e49e9322c542983d5a8374))
* **api:** update via SDK Studio ([604de63](https://github.com/ServalHQ/serval-go/commit/604de6309f801e20233d7f84ca65f550121d7132))
* **api:** update via SDK Studio ([9c699e3](https://github.com/ServalHQ/serval-go/commit/9c699e3f6e2179a3e91f9b51572db5c0d5e58e7c))
* **api:** update via SDK Studio ([2b85267](https://github.com/ServalHQ/serval-go/commit/2b852678ccc597fba801d6d64118d3bce4dcbb80))
* **api:** update via SDK Studio ([6d54be0](https://github.com/ServalHQ/serval-go/commit/6d54be0c903b78f2413b0b7942175ac9f82632c7))
* **api:** update via SDK Studio ([ff8acd1](https://github.com/ServalHQ/serval-go/commit/ff8acd1b27e7ef40b6465a64a033f1cc56896fd9))
* **api:** update via SDK Studio ([7cb216b](https://github.com/ServalHQ/serval-go/commit/7cb216bcd5b59c802abd1a9e4f7cbe93df42180b))
* **api:** update via SDK Studio ([076aa0f](https://github.com/ServalHQ/serval-go/commit/076aa0f0bd268958e8d0d4acf3a4c60689f3c390))
* **api:** update via SDK Studio ([3556d05](https://github.com/ServalHQ/serval-go/commit/3556d05666c3f56dad9bf98ccffd682938748669))
* **api:** update via SDK Studio ([00ac9dc](https://github.com/ServalHQ/serval-go/commit/00ac9dca6cf452b8c55cbf4ed1687369e8c8722f))
* **api:** update via SDK Studio ([3f9a0e2](https://github.com/ServalHQ/serval-go/commit/3f9a0e27a5f4a343ac3ed5478514ada97f0a3288))
* **api:** update via SDK Studio ([5fa96e7](https://github.com/ServalHQ/serval-go/commit/5fa96e7fac494e5b62792ced31c3172cfb56f8c5))
* **api:** update via SDK Studio ([140bc4f](https://github.com/ServalHQ/serval-go/commit/140bc4f97b50f36c45703f35ab3389016ce33db5))
* **api:** update via SDK Studio ([49e8a33](https://github.com/ServalHQ/serval-go/commit/49e8a335d1bf0eb039d08658fcd55aa67e5f6309))
* **api:** update via SDK Studio ([b1844e5](https://github.com/ServalHQ/serval-go/commit/b1844e50abec55c07899c2f9a278c89579e22aa4))


### Bug Fixes

* close body before retrying ([968809a](https://github.com/ServalHQ/serval-go/commit/968809a66a02a0698c86821337f1ce49bc79d52a))


### Chores

* update SDK settings ([df5f224](https://github.com/ServalHQ/serval-go/commit/df5f224309683d24294936686237192aff698346))

# Orion CLI

Currently this only contains the start of the project structure, a test concept, Makefile and Readme. From a terminal, cd into the `cli` folder and run:

```make build```
```make run```
```make test```

We are currently only using go-cobra and go-testing packages just to get something committed. 

# Orion CLI Commands

List help commands:

```orion --help```

Setup template in current directory:

```orion setup --template "bedrock/azure-simple"```

> The flag `--template` is specified using the following structure:
>
> `[github repo name]/[orion template name]`

Test the current template in current directory:

```orion run --docker```

# Orion Usage
1. Create project dir: my_orion_proj
2. `cd` to my_orion_proj
3. Run `orion setup --template "orion/azure-simple"`
4. Follow prompts for cloud provider info

```
ianphil@afropro.local [~/src/tmp/my_orion_proj]
> tree -a
.
├── .env
├── infra
│   └── templates
│       └── azure-simple-hw
│           ├── backend.tf
│           ├── main.tf
│           ├── outputs.tf
│           ├── provider.tf
│           ├── test
│           │   └── integration
│           │       └── azure_simple_integration_test.go
│           └── variables.tf
```
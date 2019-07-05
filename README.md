# argo-scripts

Currently there is one script `create-webhooks` that receives an input YAML file (see
[here](./values.yaml) expected values), creates GitHub webhooks for all listed repositories and
then generates an output YAML file that can be passed into the [argo-pipeline](https://github.com/dictybase-docker/kubernetes-charts/tree/master/argo-pipeline)
Helm chart.

## Command

```
NAME:
   argo-scripts - cli for scripts related to argo workflows and events

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     create-webhooks  creates new github webhooks based on given input yaml
     help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-format value  format of the logging out, either of json or text. (default: "json")
   --log-level value   log level for the application (default: "error")
   --help, -h          show help
   --version, -v       print the version
```

## Subcommand

```
NAME:
   main create-webhooks - creates new github webhooks based on given input yaml

USAGE:
   main create-webhooks [command options] [arguments...]

OPTIONS:
   --input-file value, -i value           input yaml file (default: "values.yaml")
   --output-file value, -o value          output yaml file (default: "hooks.yaml")
   --github-access-token value, -g value  github access token
```
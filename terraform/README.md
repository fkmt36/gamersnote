リヴィジョンを残しつつTerraformだけでECSをデプロイする方法
https://github.com/terraform-providers/terraform-provider-aws/issues/258#issuecomment-460981864

terraform state rm aws_ecs_task_definition.main
してから
terraform apply

## 環境変数
Terraform用の環境変数には`TF_VAR_`の接頭辞をつければ良い。例えば、`TF_VAR_NAME=hoge`とすると、Terraform内で`name=hoge`を参照できる。

## モジュールを指定する方法
`-target`オプションを使う。例えば、`terraform plan -targert=module.ecr`。

## 使用方法
1. terraform init
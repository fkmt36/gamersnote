# GamersNote

## CI・CD
mainブランチへのpushやmergeが行われた際、ECRにイメージをpushし、Terraformのapplyが行われる。pullリクエストが行われた際、Terraformのfmtチェックとplanが行われる。

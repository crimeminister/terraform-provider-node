# main.tf

provider "node" {
  version = "0.0.1"
}

data "node_package_json" "example" {
  path = "./package.json"
}

# null resources to create some changes in a plan
resource "null_resource" "resource_one" {
  triggers = {
    foo = "bar"
  }
}

resource "null_resource" "resource_two" {
  triggers = {
    foo = "bar"
  }
}

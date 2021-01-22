terraform {
  required_providers {
    teleport = {
      versions = ["0.0.1"]
      source = "gravitational.com/teleport/teleport"
    }
  }
}

provider "teleport" {
    addr = "localhost:3025"
    cert_path = "/home/nategadzhi/go/src/github.com/gravitational/teleport-plugins/docker/tmp/auth.crt"
    key_path = "/home/nategadzhi/go/src/github.com/gravitational/teleport-plugins/docker/tmp/auth.key"
    root_ca_path = "/home/nategadzhi/go/src/github.com/gravitational/teleport-plugins/docker/tmp/auth.cas"
}

resource "teleport_user" "tf_test" {
    name = "tf_test_user"
    roles = ["foo", "access-plugin"]

    trait {
      name = "logins"
      value = ["root", "foo"]
    }

    trait {
      name = "another_key"
      value = ["value", "of", "trait"]
    }
}
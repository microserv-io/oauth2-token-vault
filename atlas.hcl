data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/infrastructure/postgresql/repository",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev?search_path=public"
  migration {
    dir = "file://internal/infrastructure/postgresql/repository/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

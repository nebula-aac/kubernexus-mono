load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

load("@submodule_grpc//bazel:grpc_build_system.bzl", "grpc_proto_library")

grpc_proto_library(
    name = "metric",
    srcs = ["metric.proto"],
    use_external = True,
    well_known_protos = "@submodule_protobuf//:well_known_protos",
    deps = [
        ":common",
        "//google/api:metric",
        "//google/api:monitored_resource",
    ],
)

grpc_proto_library(
    name = "common",
    srcs = ["common.proto"],
    use_external = True,
    well_known_protos = "@submodule_protobuf//:well_known_protos",
    deps = [
        "//google/api:annotations",
        "//google/api:distribution",
    ],
)

grpc_proto_library(
    name = "metric_service",
    srcs = ["metric_service.proto"],
    use_external = True,
    well_known_protos = "@submodule_protobuf//:well_known_protos",
    deps = [
        ":common",
        ":metric",
        "//google/api:annotations",
        "//google/api:label",
        "//google/api:metric",
        "//google/api:monitored_resource",
        "//google/rpc:status",
    ],
)

proto_library(
    name = "monitoring_proto",
    srcs = [
        "common.proto",
        "group.proto",
        "group_service.proto",
        "metric.proto",
        "metric_service.proto",
    ],
    deps = [
        "//google/api:api_proto",
        "//google/monitoring/v3:v3_proto",
        "//google/rpc:rpc_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "monitoring_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "google.golang.org/genproto/googleapis/monitoring/v3",
    proto = ":monitoring_proto",
    deps = [
        "//google/api:annotations_proto",
        "//google/api:distribution_proto",
        "//google/api:metric_proto",
        "//google/api:monitored_resource_proto",
        "//google/monitoring/v3:common_proto",
        "//google/monitoring/v3:group_proto",
        "//google/monitoring/v3:metric_proto",
        "//google/rpc:status_proto",
    ],
)

go_library(
    name = "monitoring",
    embed = [":monitoring_go_proto"],
    importpath = "google.golang.org/genproto/googleapis/monitoring/v3",
)

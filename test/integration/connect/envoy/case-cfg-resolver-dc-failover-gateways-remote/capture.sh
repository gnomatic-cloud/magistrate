#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


snapshot_envoy_admin localhost:19000 s1 primary || true
snapshot_envoy_admin localhost:19001 s2 primary || true
snapshot_envoy_admin localhost:19002 s2 secondary || true
snapshot_envoy_admin localhost:19003 mesh-gateway secondary || true

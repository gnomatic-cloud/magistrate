/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

export default (collection, clickable, attribute, deletable) => () => {
  return collection('.consul-kv-list [data-test-tabular-row]', {
    name: attribute('data-test-kv', '[data-test-kv]'),
    kv: clickable('a'),
    actions: clickable('label'),
    ...deletable(),
  });
};

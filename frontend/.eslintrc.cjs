/* eslint-env node */
module.exports = {
  "root": true,
  "extends": [
    "plugin:vue/vue3-essential",
    "eslint:recommended"
  ],
  "env": {
    "vue/setup-compiler-macros": true
  },
  "rules": {
    'no-unused-vars': [2, {"args": "all", "argsIgnorePattern": "^_.*"}],
  }
}

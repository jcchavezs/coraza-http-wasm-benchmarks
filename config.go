package benchmarks

var Directives = `
Include @coraza.conf-recommended
Include @crs-setup.conf.example
Include @owasp_crs/*.conf
SecRuleEngine On
SecRuleRemoveByTag paranoia-level/2
SecRuleRemoveByTag paranoia-level/3
SecRuleRemoveByTag paranoia-level/4
SecRuleRemoveByTag paranoia-level/4
SecRuleRemoveByTag platform-windows
`

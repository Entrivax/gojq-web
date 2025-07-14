<!-- This is a JQ cheatsheet -->

## Basic filters

### Identity: `.` ([see doc](https://jqlang.org/manual/#identity))

<div class="examples">
	<runnable-example data-query="." data-json='"Hello world!"'></runnable-example>
	<runnable-example data-query="." data-json='0.123456789123456789'></runnable-example>
</div>

### Object identifier-index: `.foo`, `.foo.bar`, `."foo123"` ([see doc](https://jqlang.org/manual/#object-identifier-index))

<div class="examples">
	<runnable-example data-query=".foo" data-json='{"foo": "bar"}'></runnable-example>
	<runnable-example data-query=".foo.bar" data-json='{"foo": {"bar": "baz"}}'></runnable-example>
	<runnable-example data-query='.foo."bar"' data-json='{"foo": {"bar": "baz"}}'></runnable-example>
	<runnable-example data-query=".foo.baz" data-json='{"foo": {"bar": "baz"}}'></runnable-example>
</div>

### Optional object identifier-index: `.foo?`, `.foo.bar?` ([see doc](https://jqlang.org/manual/#optional-object-identifier-index))

<div class="examples">
	<runnable-example data-query=".foo?" data-json='{"foo": "bar"}'></runnable-example>
	<runnable-example data-query=".baz?" data-json='{"foo": "bar"}'></runnable-example>
</div>

### Object index: `.[<string>]` ([see doc](https://jqlang.org/manual/#object-index))

<div class="examples">
	<runnable-example data-query='.["foo"]' data-json='{"foo": "bar"}'></runnable-example>
	<runnable-example data-query='.["foo"].["bar"]' data-json='{"foo": {"bar": "baz"}}'></runnable-example>
	<runnable-example data-query='.["foo"]["bar"]' data-json='{"foo": {"bar": "baz"}}'></runnable-example>
	<runnable-example data-query='"foo" as $var | .[$var]' data-json='{"foo": "bar"}'></runnable-example>
</div>

### Array index: `.[<number>]` ([see doc](https://jqlang.org/manual/#array-index))

<div class="examples">
	<runnable-example data-query='.[0]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='.[-1]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='.[-2]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='1 as $var | .[$var]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
</div>

### Array and string slice: `.[<number>:<number>]` ([see doc](https://jqlang.org/manual/#array-string-slice))

<div class="examples">
	<runnable-example data-query='.[1:3]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='.[-3:-1]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='.[-2:]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='"Hello world!"[:5]' data-json='null'></runnable-example>
	<runnable-example data-query='6 as $var | "Hello world!"[$var:]' data-json='null'></runnable-example>
</div>

### Array/Object value iterator: `.[]`, `.[]?` ([see doc](https://jqlang.org/manual/#array-object-value-iterator))

<div class="examples">
	<runnable-example data-query='.[]' data-json='["foo", "bar", "baz", "qux"]'></runnable-example>
	<runnable-example data-query='.[]' data-json='{"foo": "bar", "baz": "qux"}'></runnable-example>
	<runnable-example data-query='.[]?' data-json='{"foo": "bar", "baz": "qux"}'></runnable-example>
	<runnable-example data-query='.[]?' data-json='"Hello world!"'></runnable-example>
</div>

### Comma: `,` ([see doc](https://jqlang.org/manual/#comma))

<div class="examples">
	<runnable-example data-query='."foo", ."bar"' data-json='{"foo": "baz", "bar": "qux"}'></runnable-example>
	<runnable-example data-query='.foo, .bar | length' data-json='{"foo": "baz", "bar": "hi"}'></runnable-example>
	<runnable-example data-query='["Hello", 42, true]' data-json='null'></runnable-example>
	<runnable-example data-query='.[4, 2]' data-json='["foo", "bar", "baz", "qux", "quux"]'></runnable-example>
</div>

### Pipe: `|` ([see doc](https://jqlang.org/manual/#pipe))

<div class="examples">
	<runnable-example data-query='.foo | length' data-json='{"foo": "bar"}'></runnable-example>
	<runnable-example data-query='.foo | .bar' data-json='{"foo": {"bar": "baz"}}'></runnable-example>
	<runnable-example data-query='.[] | length' data-json='["foo", "barr", "baz"]'></runnable-example>
	<runnable-example data-query='.[0] | .[1]' data-json='["fiz"]'></runnable-example>
	<runnable-example data-query='"Hello", "world!" | length' data-json='null'></runnable-example>
	<runnable-example data-query='["Hello", "world!"] | length' data-json='null'></runnable-example>
</div>

## Types and values

### Basic types ([see doc](https://jqlang.org/manual/#types-and-values))

<div class="examples">
	<runnable-example data-query='true, false, "Hello", null, 123, 42.69' data-json='null'></runnable-example>
</div>

### Arrays: `[]` ([see doc](https://jqlang.org/manual/#array-construction))

<div class="examples">
	<runnable-example data-query='[]' data-json='null'></runnable-example>
	<runnable-example data-query='["foo", .foo, .fizzbuzz[0]]' data-json='{"foo": "bar", "fizzbuzz": ["fizz", "buzz"]}'></runnable-example>
	<runnable-example data-query='.[] | [., .]' data-json='["Hello", "world"]'></runnable-example>
	<runnable-example data-query='["foo", .[], "bar"]' data-json='["Hello", "world"]'></runnable-example>
</div>

### Objects: `{}` ([see doc](https://jqlang.org/manual/#object-construction))

<div class="examples">
	<runnable-example data-query='{}' data-json='null'></runnable-example>
	<runnable-example data-query='{ "foo": .foo, "bar": .bar }' data-json='{"foo": "baz", "bar": "qux"}'></runnable-example>
	<runnable-example data-query='{ "foo": ., "bar": . }' data-json='["Hello", "world"]'></runnable-example>
</div>

## Built-in functions

### `walk(f)` ([see doc](https://jqlang.org/manual/#walk))

<div class="examples">
	<runnable-example data-query='walk(if type == "array" then map(. + "!") else . end)' data-json='{"foo": "bar", "baz": ["qux", "quux"]}'></runnable-example>
	<runnable-example data-query='walk(if type == "object" and .type? == "foo" then .value += " - test" else . end)' data-json='{
	"test": {
		"type": "foo",
		"value": "1234",
		"metadata": {
			"type": "foo",
			"value": "1234"
		}
	},
	"test2": {
		"type": "foo",
		"value": "12345"
	},
	"test3": {
		"type": "bar",
		"value": "22222",
		"metadata": {
			"type": "foo",
			"value": "1234",
			"metadata": {
				"type": "bar",
				"value": "1234"
			}
		}
	}
}'></runnable-example>
</div>

## Advanced examples

### Embed in items

<div class="examples">
	<runnable-example data-query=". as $data | {
	items: [.items[] | . += 
	{ 
		resources: (.resourceIds | map(
			. as $rid | 
			($data.resources[] | select(.id == $rid))
		)) 
	} | del(.resourceIds)]
}" data-json='{
	"resources": [
		{ "id": "123", "url": "http://example.com/res.png" }
	],
	"items": [
		{ "id": "1", "resourceIds": ["123"], "title": "Test" }
	]
}'></runnable-example>
</div>

### Create atlas array from items children

<div class="examples">
	<runnable-example data-query='. += {
	resources: [.items[].resources[]] | unique_by(.id),
	items: [.items[] | . + { resourceIds: [.resources[].id] } | del(.resources)]
}' data-json='{
	"items": [
		{
			"id": "1",
			"resources": [{ "id": "123", "url": "http://example.com/res.png" }],
			"title": "Test"
		}
	]
}'></runnable-example>
</div>

### Create atlas dictionary from items children

<div class="examples">
	<runnable-example data-query='. += {
	resources: [.items[].resources[]] | map({ (.id): . }) | add,
	items: [.items[] | . + { resourceIds: [.resources[].id] } | del(.resources)]
}' data-json='{
	"items": [
		{
			"id": "1",
			"resources": [{ "id": "123", "url": "http://example.com/res.png" }],
			"title": "Test"
		}
	]
}'></runnable-example>
	<runnable-example data-query='. += {
	resources: [.items[].resources[]] | map({ key: .id, value: . }) | from_entries,
	items: [.items[] | . + { resourceIds: [.resources[].id] } | del(.resources)]
}' data-json='{
	"items": [
		{
			"id": "1",
			"resources": [{ "id": "123", "url": "http://example.com/res.png" }],
			"title": "Test"
		}
	]
}'></runnable-example>
</div>
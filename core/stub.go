package core

const (
	stubHeader = `// TEMPORARY AUTOGENERATED FILE: perezajson stub code to make the package
// compilable during generation.

package `
	n = byte('\n')
)

// Static allocate
func Stub(packageName string, types []string) []byte {
	content := make([]byte, 0, getStubSize(packageName, types))

	content = append(content, stubHeader...)
	content = append(content, packageName...)

	content = append(content, n, n)

	for _, t := range types {
		content = append(content, "func ("...)
		content = append(content, t...)
		content = append(content, ") MarshalJSON() ([]byte, error) { return nil, nil }\n"...)
		content = append(content, "func (*"...)
		content = append(content, t...)
		content = append(content, ") UnmarshalJSON([]byte) error  { return nil }\n"...)
		content = append(content, n)
		content = append(content, "type PerezaJSON_exporter_"...)
		content = append(content, t...)
		content = append(content, ' ', '*')
		content = append(content, t...)
		content = append(content, n)
	}

	return content
}

func getStubSize(packageName string, types []string) int {
	//	body := `func () MarshalJSON() ([]byte, error) { return nil, nil }
	//func (*) UnmarshalJSON([]byte) error  { return nil }
	//`

	const (
		headerSize          = len(stubHeader)
		headerNextSpaceSize = 2
		bodySize            = 140 // len(body)
	)

	return headerSize +
		len(packageName) +
		headerNextSpaceSize +
		bodySize*len(types) +
		stringSliceSize(types)*4 // 4 is func: MarshalJSON & UnmarshalJSON & type alias
}

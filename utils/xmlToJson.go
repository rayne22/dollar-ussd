package utils

import "github.com/beevik/etree"

func ConvertXMLToMap(xmlString string) (map[string]interface{}, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(xmlString)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	root := doc.Root()

	// Recursively parse the XML tree
	ParseXMLNode(root, result)

	return result, nil
}

func ParseXMLNode(node *etree.Element, result map[string]interface{}) {
	// Check if the node has child elements
	if len(node.ChildElements()) == 0 {
		// Leaf node, store the value
		result[node.Tag] = node.Text()
	} else {
		// Non-leaf node, recursively parse child elements
		childResult := make(map[string]interface{})
		for _, child := range node.ChildElements() {
			ParseXMLNode(child, childResult)
		}
		result[node.Tag] = childResult
	}
}

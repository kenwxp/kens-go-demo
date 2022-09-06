package region

import (
	"encoding/json"
	"kens/demo/cache"
	"kens/demo/storage/types"
	"strconv"
	"strings"
)

type Node struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
	Code  string `json:"code"`
	Child []Node `json:"child"`
}

func genNode(dataList []types.Region, currentNode Node, recursion int) ([]Node, error) {
	currentNodeLevel, _ := strconv.Atoi(currentNode.Level)
	child := make([]Node, 0)
	if currentNodeLevel >= recursion {
		return child, nil
	}
	for _, item := range dataList {
		if item.RegionParentId == currentNode.Id {
			node := Node{
				Id:    item.RegionId,
				Name:  item.RegionName,
				Level: item.RegionLevel,
				Code:  item.RegionCode,
			}
			childa, err := genNode(dataList, node, recursion)
			if err != nil {
				return nil, err
			}
			node.Child = childa
			child = append(child, node)
		}

	}
	return child, nil
}

func GetRegions(search string, searchType SearchType, searchLevel SearchLevel) ([]Node, error) {
	list := make([]Node, 0)
	if searchLevel == SearchSelf {
		node, err := GetSelfRegion(search, searchType)
		if err != nil {
			return nil, err
		}
		if node != nil {
			list = append(list, *node)
		}
	} else if searchLevel == SearchParent {
		node, err := GetParentRegion(search, searchType)
		if err != nil {
			return nil, err
		}
		if node != nil {
			list = append(list, *node)
		}

	} else if searchLevel == SearchChild {
		nodes, err := GetChildRegion(search, searchType)
		if err != nil {
			return nil, err
		}
		list = nodes
	}
	return list, nil
}
func GetChildRegion(search string, searchType SearchType) ([]Node, error) {
	return getChildNodes(getRegionTreeFromCache(), search, searchType)
}

func getChildNodes(regionTree *Node, search string, searchType SearchType) ([]Node, error) {
	childList := regionTree.Child
	if (searchType == SearchById && "0000000" == search) ||
		(searchType == SearchByCode && "0" == search) ||
		(searchType == SearchByName && strings.Contains("中国", search)) {
		return childList, nil
	}
	for _, item := range childList {
		if (searchType == SearchById && item.Id == search) ||
			(searchType == SearchByCode && item.Code == search) ||
			(searchType == SearchByName && strings.Contains(item.Name, search)) {
			return item.Child, nil
		} else {
			childNodes, err := getChildNodes(&item, search, searchType)
			if err != nil {
				return nil, err
			}
			if childNodes != nil {
				return childNodes, nil
			}
		}
	}
	return nil, nil
}

func GetSelfRegion(search string, searchType SearchType) (*Node, error) {
	return getNodes(getRegionTreeFromCache(), search, searchType)
}
func getNodes(regionTree *Node, search string, searchType SearchType) (*Node, error) {
	childList := regionTree.Child
	if (searchType == SearchById && "0000000" == search) ||
		(searchType == SearchByCode && "0" == search) ||
		(searchType == SearchByName && strings.Contains("中国", search)) {
		return regionTree, nil
	}
	for _, item := range childList {
		if (searchType == SearchById && item.Id == search) ||
			(searchType == SearchByCode && item.Code == search) ||
			(searchType == SearchByName && strings.Contains(item.Name, search)) {
			return &item, nil
		} else {
			node, err := getNodes(&item, search, searchType)
			if err != nil {
				return nil, err
			}
			if node != nil {
				return node, nil
			}
		}
	}
	return nil, nil
}

func GetParentRegion(search string, searchType SearchType) (*Node, error) {
	return getParentNode(getRegionTreeFromCache(), search, searchType)
}
func getParentNode(regionTree *Node, search string, searchType SearchType) (*Node, error) {
	childList := regionTree.Child
	if (searchType == SearchById && "0000000" == search) ||
		(searchType == SearchByCode && "0" == search) ||
		(searchType == SearchByName && strings.Contains("中国", search)) {
		return nil, nil
	}
	for _, item := range childList {
		if (searchType == SearchById && item.Id == search) ||
			(searchType == SearchByCode && item.Code == search) ||
			(searchType == SearchByName && strings.Contains(item.Name, search)) {
			return regionTree, nil
		} else {
			parentNode, err := getParentNode(&item, search, searchType)
			if err != nil {
				return nil, err
			}
			if parentNode != nil {
				return parentNode, nil
			}
		}
	}
	return nil, nil
}
func getRegionTreeFromCache() *Node {
	regionTreeRaw, isExist := cache.CacheAll.Get(CacheRegionNodeKey)
	if isExist {
		bytes, err := json.Marshal(regionTreeRaw)
		if err != nil {
			return nil
		}
		regionTree := &Node{}
		json.Unmarshal(bytes, regionTree)
		return regionTree
	}
	return nil
}

func GetRegionNameByCode(key string) string {
	node, err := GetSelfRegion(key, SearchByCode)
	retStr := ""
	if err == nil && node != nil {
		retStr = node.Name
		if node.Level != "1" {
			return getParentRegionNameByCode(node.Code) + retStr
		}

	}
	return retStr
}
func getParentRegionNameByCode(key string) string {
	node, err := GetParentRegion(key, SearchByCode)
	retStr := ""
	if err == nil && node != nil {
		retStr = node.Name
		if node.Level != "1" {
			return getParentRegionNameByCode(node.Code) + retStr
		}
	}
	return retStr
}

func GetSearchRegionCode(code string) string {
	if code == "" || len(code) != 6 {
		return ""
	}
	one := code[0:2]
	two := code[2:4]
	three := code[4:]
	search := ""
	if three == "00" {
		if two == "00" {
			if one == "00" {
				search = ""
			} else {
				search = one
			}
		} else {
			search = one + two
		}
	} else {
		search = one + two + three
	}
	return search
}

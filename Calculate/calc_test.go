package Calculate

import (
	"testing"
)

func TestOrderinventory(t *testing.T) {
	var items []string
	availableItems := []string{"沐浴乳", "洗髮精", "酒精"} // 假設這些物品有庫存
	unavailableItems := []string{"咖啡包", "餅乾"}      // 假設這些物品沒有庫存

	// 測試有庫存的物品
	for _, item := range availableItems {
		Orderinventory(item, "3", &items)
		if len(items) > 0 {
			t.Errorf("Expected item %s not to be ordered, but it was not", item)
		}
	}

	// 測試沒有庫存的物品
	for _, item := range unavailableItems {
		Orderinventory(item, "0", &items)
		if len(items) == 0 {
			t.Errorf("Expected item %s to be ordered, but it was", item)
		}
	}
}

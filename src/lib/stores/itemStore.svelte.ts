import type { DofusItem } from "$lib/types/dofus-items";

class ItemStore {
    private selectedItems: DofusItem[] = $state([]);

    constructor() { }

    getSelectedItems(): DofusItem[] {
        return this.selectedItems;
    }

    selectItem(item: DofusItem) {
        if (this.selectedItems.includes(item)) return;

        this.selectedItems.push(item);
    }
}

const itemsStore = new ItemStore();

export default itemsStore;
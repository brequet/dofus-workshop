export interface Recipe {
	itemId: number;
	quantity: number;
}

export interface DofusItem {
	id: number;
	name: string;
	iconId: number;
	recipe: Recipe[];
	selectedCount: number;
}

export interface DofusItems {
	[key: string]: DofusItem;
}

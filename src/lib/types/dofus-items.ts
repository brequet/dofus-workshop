export interface Recipe {
    itemId: number;
    quantity: number;
}

export interface DofusItem {
    id: number;
    name: string;
    iconId: number;
    recipe: Recipe[];
}

export interface DofusItems {
    [key: string]: DofusItem;
}

import type { DofusItem, DofusItems } from '$lib/types/dofus-items';
import { readFile } from 'fs/promises';

let cachedItems: DofusItems | null = null;

function getItems(): DofusItems {
    if (!cachedItems) {
        throw new Error('Dofus items not loaded');
    }
    return cachedItems;
}

export async function loadDofusItems(): Promise<DofusItems> {
    if (cachedItems) {
        return cachedItems;
    }

    try {
        const data = await readFile('./static/dofus/dofus-items.json', 'utf-8');
        cachedItems = JSON.parse(data) as DofusItems;
        return cachedItems;
    } catch (error) {
        console.error('Error loading Dofus items:', error);
        throw new Error('Failed to load Dofus items');
    }
}

export function getDofusItem(id: string): DofusItem | undefined {
    if (!cachedItems) {
        throw new Error('Dofus items not loaded');
    }
    return getItems()[id];
}

export function searchDofusItems(search: string): DofusItem[] {
    if (!cachedItems) {
        throw new Error('Dofus items not loaded');
    }

    if (search.length < 3) {
        return [];
    }

    const items = getItems();
    return Object.values(items)
        .filter((item) => item.recipe.length > 0)
        .filter((item) => item.name.toLowerCase().includes(search.toLowerCase()))
        .sort((a, b) => a.name.localeCompare(b.name))
        .slice(0, 25);
}

import { loadDofusItems, searchDofusItems } from '$lib/server/dofusItemsLoader';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ params }) => {
    await loadDofusItems();
    const item = searchDofusItems(params.input);

    return json(item);
};

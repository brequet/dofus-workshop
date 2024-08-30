import { getDofusItem, loadDofusItems } from '$lib/server/dofusItemsLoader';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ params }) => {
    await loadDofusItems();
    const item = getDofusItem(params.id);

    if (item) {
        return json(item);
    } else {
        return new Response('Item not found', { status: 404 });
    }
};

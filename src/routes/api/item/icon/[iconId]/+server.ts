import { error } from '@sveltejs/kit';

export async function GET({ params, fetch, setHeaders }) {
    const { iconId } = params;
    const imageUrl = `https://api.dofusdb.fr/img/items/250/${iconId}.png`;

    try {
        const response = await fetch(imageUrl);

        if (!response.ok) {
            throw error(response.status, 'Failed to fetch image');
        }

        const imageBuffer = await response.arrayBuffer();
        const contentType = response.headers.get('content-type');

        // Set caching headers
        setHeaders({
            'Cache-Control': 'public, max-age=3600', // Cache for 1 hour
            'Content-Type': contentType
        });

        return new Response(imageBuffer, {
            headers: {
                'Content-Type': contentType
            }
        });
    } catch (err) {
        console.error(`Error fetching image '${imageUrl}' :`, err);
        throw error(500, 'Internal Server Error');
    }
}
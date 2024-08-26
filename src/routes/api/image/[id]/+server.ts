import { error } from '@sveltejs/kit';
import fs from 'fs';
import path from 'path';

export function GET({ params }) {
	const { id } = params;
	const imagePath = path.join(process.cwd(), 'static', 'dofus', 'icons', id + '.png');

	if (fs.existsSync(imagePath)) {
		const image = fs.readFileSync(imagePath);
		return new Response(image, {
			headers: {
				'Content-Type': 'image/png', // Adjust based on your image type
				'Cache-Control': 'public, max-age=604800' // Example caching header
			}
		});
	}

	throw error(404, 'Image not found');
}

import { useEffect, useState } from "react";
import { getPosts } from "../../api/posts";
import type { post } from "../../api/types";
import Post from "./post";

function PostList() {
	const [loading, setLoading] = useState(true);
	const [posts, setPosts] = useState<post[]>([]);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const posts = await getPosts();
				setPosts(posts.posts);
			} catch (error) {
				console.error("failed to fetch posts", error);
			} finally {
				setLoading(false);
			}
		};
		fetchData();
	}, []);

	if (loading) {
		return <div>loading...</div>;
	}

	return (
		<div className="grid grid-cols-2 md:grid-cols-4 gap-x-3 gap-y-2 mt-5 p-3">
			{posts.map((post) => {
				return <Post key={post.id} post={post} />;
			})}
		</div>
	);
}

export default PostList;

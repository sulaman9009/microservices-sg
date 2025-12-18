import type { queryPost } from "../../api/types";
import Comments from "../comments/comments";
import CreateCommentForm from "../comments/create-comment-form";

type props = {
	post: queryPost;
};

function Post({ post }: props) {
	return (
		<div className="border border-zinc-200 p-2 rounded-md shadow-md">
			<p>{post.title}</p>
			<hr className="my-3" />
			<CreateCommentForm postId={post.id} />
			<hr className="my-3" />
			<p className="mb-2">Comments:</p>
			<Comments comments={post.comments} />
		</div>
	);
}

export default Post;

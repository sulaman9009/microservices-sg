import { useActionState } from "react";
import { createComment } from "../../api/comments";

type props = {
	postId: string;
};

function CreateCommentForm({ postId }: props) {
	const submitForm = async (_: unknown, data: FormData) => {
		const content = data.get("content")?.toString();
		if (!content) {
			console.error("content not set");
			return { result: "failure" };
		}
		try {
			await createComment(postId, content);
		} catch (error) {
			console.error("failed to create comment:", error);
			return { result: "failure" };
		}
		return { result: "success" };
	};

	const [_, action, isLoading] = useActionState(submitForm, undefined);
	return (
		<div>
			<form action={action}>
				<div className="flex flex-col">
					<label htmlFor="content" className="text-zinc-800 text-md mb-1">
						New Comment
					</label>
					<input
						type="text"
						name="content"
						id="content"
						className="border-zinc-400 rounded-md border p-1"
					/>
				</div>
				<button
					type="submit"
					className="bg-sky-400 text-zinc-800 py-1 px-2 mt-3 rounded-md hover:bg-blue-500 transition-colors ease-in-out cursor-pointer"
				>
					{isLoading ? "submitting..." : "submit"}
				</button>
			</form>
		</div>
	);
}

export default CreateCommentForm;

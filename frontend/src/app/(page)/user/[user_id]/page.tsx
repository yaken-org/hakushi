export default function PostDetailPage({
    params: {
        user_id
    }
}: {
    params: {
        user_id: string;
    };
}) {
    return (
        <div>
            <h1>Post Detail</h1>
            <p>Post ID: {user_id}</p>
        </div>
    )
}
const get_user_account_data = async (sub: string) => {
    try {
        const user_account_data_response = await fetch(process.env.BACKEND_API_ENDPOINT + "/user_account/" + sub);

        if (!user_account_data_response.ok) {
            return { is_success: false, message: "Failed to fetch user account data" };
        }
        return { is_success: true, data: await user_account_data_response.json() };
    } catch (error) {
        return { is_success: false, message: "Failed to fetch user account data" };
    }
}

export { get_user_account_data };
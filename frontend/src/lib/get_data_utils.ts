const backend_url = process.env.BACKEND_API_ENDPOINT ?? "http://hakushi-backend.default.svc.cluster.local/api";

const get_account_data_by_sub = async (sub: string) => {
    try {
        const account_data_response = await fetch(`${backend_url}/account/sub/${sub}`);
        console.log(account_data_response);

        if (!account_data_response.ok && account_data_response.status !== 404) {
            return { is_success: false, message: "Failed to fetch user account data" };
        } else if (account_data_response.status === 404) {
            return { is_success: true, status_code: account_data_response.status };
        }
        return { is_success: true, status_code: account_data_response.status, data: await account_data_response.json() };
    } catch (error) {
        return { is_success: false, message: "Failed to fetch user account data", data: error };
    }
}

const get_account_data_by_name = async (name: string) => {
    try {
        const account_data_response = await fetch(process.env.BACKEND_API_ENDPOINT + "/account/name/" + name);
        console.log(account_data_response);

        if (!account_data_response.ok && account_data_response.status !== 404) {
            return { is_success: false, message: "Failed to fetch user account data" };
        } else if (account_data_response.status === 404) {
            return { is_success: true, status_code: account_data_response.status };
        }
        return { is_success: true, status_code: account_data_response.status, data: await account_data_response.json() };
    } catch (error) {
        return { is_success: false, message: "Failed to fetch user account data", data: error };
    }
}

export { get_account_data_by_sub, get_account_data_by_name };
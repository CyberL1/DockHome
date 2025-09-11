import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
  const dataReq = await fetch("/api/domain");
  return await dataReq.json();
};

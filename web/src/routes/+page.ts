import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
  const dataReq = await fetch("/api/data");
  return await dataReq.json();
};

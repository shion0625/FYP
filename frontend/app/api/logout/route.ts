import { deleteToken } from '@/utils/cookie';

export async function GET() {
  deleteToken();
  return Response.json('success');
}

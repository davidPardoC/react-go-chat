import { useQuery } from "@tanstack/react-query";
import { getAllUsers } from "../services/users.service";
import { Input } from "@/components/ui/input";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { getTokenPayload } from "@/utils/auth";
import { Link } from "wouter";

const UsersList = () => {
  const { data } = useQuery({
    queryKey: ["users"],
    queryFn: getAllUsers,
  });

  return (
    <>
      <Input placeholder="Search.." />
      <ul className="mt-2">
        {data?.map((user) => (
          <Link key={user.id} href={`/chat?recipient_id=${user.id}`}>
            <li className="flex items-center gap-2 py-2">
              <Avatar>
                <AvatarFallback>
                  {user.username.toUpperCase().slice(0, 2)}
                </AvatarFallback>
              </Avatar>
              {user.username}{" "}
              {getTokenPayload().username == user.username && <span>(Me)</span>}
            </li>
          </Link>
        ))}
      </ul>
    </>
  );
};

export default UsersList;

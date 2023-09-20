import { useMutation } from "@tanstack/react-query";
import axios from "axios";

interface CreateObjectRequest {
  file: File;
}

export const useCreateObject = () => {
  const mutation = useMutation({
    mutationFn: (req: CreateObjectRequest) => {
      const path = process.env.NEXT_PUBLIC_OBJ_URL_PREFIX + '/obj/api/v1/wastes';
      return axios.post(path, req.file, {
        headers: {
          'Content-Type': req.file.type,
        }
      });
    },
  });

  return mutation;
};
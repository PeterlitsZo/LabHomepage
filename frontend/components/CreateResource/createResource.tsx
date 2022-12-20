import { useEffect, useState } from "react";

import Input from "../Input";
import Textarea from "../Textarea";
import Button from "../Button";
import { useCreateResource } from "../../requests/resources";
import { useRouter } from "next/router";

export const CreateResource = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const createResource = useCreateResource();
  const router = useRouter();

  useEffect(() => {
    if (createResource.isSuccess) {
      router.push("/resources");
    }
  }, [createResource, router]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <Input
        placeholder="Input the title of your new resource here"
        className="mb-4"
        onChange={(e) => setTitle(e.target.value)}
      />
      <Textarea
        rows={20}
        placeholder="Input the content of your new resource here"
        onChange={(e) => setContent(e.target.value)}
      />
      <Button className="mt-4" onClick={() => {
        createResource.mutate({ title, content });
      }}>
        Publish
      </Button>
    </main>
  )
}
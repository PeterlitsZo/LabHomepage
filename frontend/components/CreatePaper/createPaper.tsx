import { useEffect, useState } from "react";

import Input from "../Input";
import Textarea from "../Textarea";
import Button from "../Button";
import { useCreatePaper } from "../../requests/papers";
import { useRouter } from "next/router";

export const CreatePaper = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const createPaper = useCreatePaper();
  const router = useRouter();

  useEffect(() => {
    if (createPaper.isSuccess) {
      router.push("/papers");
    }
  }, [createPaper, router]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <Input
        placeholder="Input the title of your new paper here"
        className="mb-4"
        onChange={(e) => setTitle(e.target.value)}
      />
      <Textarea
        rows={20}
        placeholder="Input the content of your new paper here"
        onChange={(e) => setContent(e.target.value)}
      />
      <Button className="mt-4" onClick={() => {
        createPaper.mutate({ title, content });
      }}>
        Publish
      </Button>
    </main>
  )
}
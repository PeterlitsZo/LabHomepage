import { useEffect, useState } from "react";

import Input from "../Input";
import Textarea from "../Textarea";
import Button from "../Button";
import { useCreatePerson } from "../../requests/people";
import { useRouter } from "next/router";

export const CreatePerson = () => {
  const [name, setName] = useState("");
  const [content, setContent] = useState("");

  const createPerson = useCreatePerson();
  const router = useRouter();

  useEffect(() => {
    if (createPerson.isSuccess) {
      router.push("/people");
    }
  }, [createPerson, router]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <Input
        placeholder="Input the title of your new person here"
        className="mb-4"
        onChange={(e) => setName(e.target.value)}
      />
      <Textarea
        rows={20}
        placeholder="Input the content of your new person here"
        onChange={(e) => setContent(e.target.value)}
      />
      <Button className="mt-4" onClick={() => {
        createPerson.mutate({ name, content });
      }}>
        Publish
      </Button>
    </main>
  )
}
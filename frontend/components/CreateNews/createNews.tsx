import { useEffect, useState } from "react";

import Input from "../Input";
import Textarea from "../Textarea";
import Button from "../Button";
import { useCreateNews } from "../../requests/news";
import { useRouter } from "next/router";

export const CreateNews = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const createNews = useCreateNews();
  const router = useRouter();

  useEffect(() => {
    if (createNews.isSuccess) {
      router.push("/news");
    }
  }, [createNews, router]);

  return (
    <main className="max-w-3xl m-auto py-24">
      <Input
        placeholder="Input the title of your new news here"
        className="mb-4"
        onChange={(e) => setTitle(e.target.value)}
      />
      <Textarea
        rows={20}
        placeholder="Input the content of your new news here"
        onChange={(e) => setContent(e.target.value)}
      />
      <Button className="mt-4" onClick={() => {
        createNews.mutate({ title, content });
      }}>
        Publish
      </Button>
    </main>
  )
}
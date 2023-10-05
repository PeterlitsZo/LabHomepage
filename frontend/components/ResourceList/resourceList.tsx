import { useDeleteResource, useListResource } from "../../requests/resources"
import List from "../List";

export const ResourceList = () => {
  const listResourceQuery = useListResource();
  const deleteResourceMutation = useDeleteResource();

  return (
    <List
      createItem={{
        url: '/resources/create',
        label: 'Create New Resource',
      }}
      items={{
        whenRender: listResourceQuery.isSuccess,
        refPrefix: '/resources/',
        deleteItem(id) {
          deleteResourceMutation.mutate(id);
        },
        getTitle: d => d.title,
        data: listResourceQuery.data?.data.resources,
      }}
    />
  )
}
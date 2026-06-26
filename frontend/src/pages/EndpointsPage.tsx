import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Plus, Pencil, Trash2, Search } from "lucide-react";
import {
  Badge,
  Button,
  Card,
  EmptyState,
  ErrorState,
  Input,
  Modal,
  Select,
  Table,
  TableSkeleton,
  Typography,
} from "@/components/ui";
import {
  useCreateEndpoint,
  useDeleteEndpoint,
  useEndpoints,
  useUpdateEndpoint,
} from "@/hooks/useEndpoints";
import { formatDate, formatMs, getStatusColor } from "@/utils/format";
import type { Endpoint, EndpointCreateUpdate } from "@/types/api";

const emptyForm: EndpointCreateUpdate = {
  name: "",
  url: "",
  expected_status: 200,
};

export function EndpointsPage() {
  const navigate = useNavigate();
  const { data, isLoading, isError, refetch } = useEndpoints();
  const createMutation = useCreateEndpoint();
  const updateMutation = useUpdateEndpoint();
  const deleteMutation = useDeleteEndpoint();

  const [search, setSearch] = useState("");
  const [statusFilter, setStatusFilter] = useState("all");
  const [modalOpen, setModalOpen] = useState(false);
  const [editing, setEditing] = useState<Endpoint | null>(null);
  const [form, setForm] = useState<EndpointCreateUpdate>(emptyForm);
  const [deleteConfirm, setDeleteConfirm] = useState<Endpoint | null>(null);

  const filtered = (data ?? []).filter((ep) => {
    const matchesSearch =
      ep.name.toLowerCase().includes(search.toLowerCase()) ||
      ep.url.toLowerCase().includes(search.toLowerCase());
    const matchesStatus = statusFilter === "all" || ep.status === statusFilter;
    return matchesSearch && matchesStatus;
  });

  const openCreate = () => {
    setEditing(null);
    setForm(emptyForm);
    setModalOpen(true);
  };

  const openEdit = (ep: Endpoint) => {
    setEditing(ep);
    setForm({
      name: ep.name,
      url: ep.url,
      expected_status: ep.expected_status,
    });
    setModalOpen(true);
  };

  const handleSave = async () => {
    try {
      if (editing) {
        await updateMutation.mutateAsync({ id: editing.id, payload: form });
      } else {
        await createMutation.mutateAsync(form);
      }
      setModalOpen(false);
      refetch();
    } catch {
      // TODO(API): Remove mock success when POST/PUT /endpoints is connected
      setModalOpen(false);
      refetch();
    }
  };

  const handleDelete = async () => {
    if (!deleteConfirm) return;
    try {
      await deleteMutation.mutateAsync(deleteConfirm.id);
    } catch {
      // TODO(API): Remove mock success when DELETE /endpoints/{id} is connected
    }
    setDeleteConfirm(null);
    refetch();
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <Typography variant="h5" color="white" fontWeight="bold">
            Endpoints
          </Typography>
          <Typography variant="body2" color="text">
            Manage and monitor your API endpoints
          </Typography>
        </div>
        <Button onClick={openCreate}>
          <Plus className="h-4 w-4" />
          Add Endpoint
        </Button>
      </div>

      <Card>
        <div className="mb-4 flex flex-col gap-3 sm:flex-row">
          <div className="flex-1">
            <Input
              placeholder="Search endpoints..."
              icon={<Search className="h-4 w-4" />}
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
          </div>
          <Select
            options={[
              { value: "all", label: "All Statuses" },
              { value: "healthy", label: "Healthy" },
              { value: "unhealthy", label: "Unhealthy" },
            ]}
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value)}
            className="sm:w-48"
          />
        </div>

        {isLoading && <TableSkeleton />}
        {isError && <ErrorState onRetry={() => refetch()} />}
        {!isLoading && !isError && filtered.length === 0 && (
          <EmptyState
            title="No endpoints found"
            description="Create your first endpoint to start monitoring."
            action={
              <Button onClick={openCreate}>
                <Plus className="h-4 w-4" />
                Add Endpoint
              </Button>
            }
          />
        )}
        {!isLoading && !isError && filtered.length > 0 && (
          <Table
            data={filtered}
            keyExtractor={(row) => row.id}
            onRowClick={(row) => navigate(`/endpoints/${row.id}`)}
            columns={[
              {
                key: "name",
                header: "Name",
                render: (r) => <span className="font-medium">{r.name}</span>,
              },
              {
                key: "url",
                header: "URL",
                render: (r) => (
                  <span className="text-text truncate max-w-[200px] block">
                    {r.url}
                  </span>
                ),
              },
              {
                key: "status",
                header: "Status",
                render: (r) => (
                  <Badge color={getStatusColor(r.status ?? "unknown")}>
                    {r.status ?? "unknown"}
                  </Badge>
                ),
              },
              { key: "expected_status", header: "Expected Status" },
              {
                key: "last_checked",
                header: "Last Checked",
                render: (r) =>
                  r.last_checked ? formatDate(r.last_checked) : "—",
              },
              {
                key: "response_time",
                header: "Response Time",
                render: (r) =>
                  r.response_time ? formatMs(r.response_time) : "—",
              },
              {
                key: "actions",
                header: "Actions",
                render: (r) => (
                  <div
                    className="flex gap-1"
                    onClick={(e) => e.stopPropagation()}
                  >
                    <Button
                      variant="text"
                      iconOnly
                      size="small"
                      onClick={() => openEdit(r)}
                    >
                      <Pencil className="h-4 w-4" />
                    </Button>
                    <Button
                      variant="text"
                      iconOnly
                      size="small"
                      onClick={() => setDeleteConfirm(r)}
                    >
                      <Trash2 className="h-4 w-4 text-error" />
                    </Button>
                  </div>
                ),
              },
            ]}
          />
        )}
      </Card>

      <Modal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        title={editing ? "Edit Endpoint" : "Create Endpoint"}
        footer={
          <>
            <Button
              variant="outlined"
              color="white"
              onClick={() => setModalOpen(false)}
            >
              Cancel
            </Button>
            <Button
              onClick={handleSave}
              loading={createMutation.isPending || updateMutation.isPending}
            >
              {editing ? "Save Changes" : "Create"}
            </Button>
          </>
        }
      >
        <div className="space-y-4">
          <div>
            <Typography variant="button" color="white" className="mb-1 block">
              Name
            </Typography>
            <Input
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              placeholder="Auth API"
            />
          </div>
          <div>
            <Typography variant="button" color="white" className="mb-1 block">
              URL
            </Typography>
            <Input
              value={form.url}
              onChange={(e) => setForm({ ...form, url: e.target.value })}
              placeholder="https://api.example.com/health"
            />
          </div>
          <div>
            <Typography variant="button" color="white" className="mb-1 block">
              Expected Status Code
            </Typography>
            <Input
              type="number"
              value={form.expected_status}
              onChange={(e) =>
                setForm({ ...form, expected_status: Number(e.target.value) })
              }
            />
          </div>
        </div>
      </Modal>

      <Modal
        open={!!deleteConfirm}
        onClose={() => setDeleteConfirm(null)}
        title="Delete Endpoint"
        footer={
          <>
            <Button
              variant="outlined"
              color="white"
              onClick={() => setDeleteConfirm(null)}
            >
              Cancel
            </Button>
            <Button
              color="error"
              onClick={handleDelete}
              loading={deleteMutation.isPending}
            >
              Delete
            </Button>
          </>
        }
      >
        <Typography variant="body2" color="text">
          Are you sure you want to delete{" "}
          <strong className="text-text-focus">{deleteConfirm?.name}</strong>?
          This action cannot be undone.
        </Typography>
      </Modal>
    </div>
  );
}

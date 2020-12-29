// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/playlist"
	"github.com/tanapon395/playlist-video/ent/playlist_video"
	"github.com/tanapon395/playlist-video/ent/user"
)

// PlaylistCreate is the builder for creating a Playlist entity.
type PlaylistCreate struct {
	config
	mutation *PlaylistMutation
	hooks    []Hook
}

// SetTitle sets the title field.
func (pc *PlaylistCreate) SetTitle(s string) *PlaylistCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetOwnerID sets the owner edge to User by id.
func (pc *PlaylistCreate) SetOwnerID(id int) *PlaylistCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pc *PlaylistCreate) SetNillableOwnerID(id *int) *PlaylistCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the owner edge to User.
func (pc *PlaylistCreate) SetOwner(u *User) *PlaylistCreate {
	return pc.SetOwnerID(u.ID)
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (pc *PlaylistCreate) AddPlaylistVideoIDs(ids ...int) *PlaylistCreate {
	pc.mutation.AddPlaylistVideoIDs(ids...)
	return pc
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (pc *PlaylistCreate) AddPlaylistVideos(p ...*Playlist_Video) *PlaylistCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pc *PlaylistCreate) Mutation() *PlaylistMutation {
	return pc.mutation
}

// Save creates the Playlist in the database.
func (pc *PlaylistCreate) Save(ctx context.Context) (*Playlist, error) {
	if err := pc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Playlist
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaylistCreate) SaveX(ctx context.Context) *Playlist {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PlaylistCreate) preSave() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := playlist.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	return nil
}

func (pc *PlaylistCreate) sqlSave(ctx context.Context) (*Playlist, error) {
	pl, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pl.ID = int(id)
	return pl, nil
}

func (pc *PlaylistCreate) createSpec() (*Playlist, *sqlgraph.CreateSpec) {
	var (
		pl    = &Playlist{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playlist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldTitle,
		})
		pl.Title = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.OwnerTable,
			Columns: []string{playlist.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.PlaylistVideosTable,
			Columns: []string{playlist.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pl, _spec
}

// PlaylistCreateBulk is the builder for creating a bulk of Playlist entities.
type PlaylistCreateBulk struct {
	config
	builders []*PlaylistCreate
}

// Save creates the Playlist entities in the database.
func (pcb *PlaylistCreateBulk) Save(ctx context.Context) ([]*Playlist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Playlist, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*PlaylistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *PlaylistCreateBulk) SaveX(ctx context.Context) []*Playlist {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
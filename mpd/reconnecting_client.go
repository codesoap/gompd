// Copyright 2009 The GoMPD Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package mpd provides the client side interface to MPD (Music Player Daemon).
// The protocol reference can be found at http://www.musicpd.org/doc/protocol/index.html
package mpd

type ReconnectingClient struct {
	client   *Client
	network  string
	address  string
	password *string
}

func NewReconnectingClient(network, addr string) (*ReconnectingClient, error) {
	c, err := Dial(network, addr)
	return &ReconnectingClient{c, network, addr, nil}, err
}

func NewReconnectingClientAuthenticated(network, addr, password string) (*ReconnectingClient, error) {
	c, err := DialAuthenticated(network, addr, password)
	return &ReconnectingClient{c, network, addr, &password}, err
}

func (r *ReconnectingClient) reconnect() (err error) {
	if r.password == nil {
		r.client, err = Dial(r.network, r.address)
	} else {
		r.client, err = DialAuthenticated(r.network, r.address, *r.password)
	}
	return err
}

// Add does the same as the Add method of Client, but tries to reconnect
// once if an error occurs.
func (r *ReconnectingClient) Add(uri string) error {
	err := r.client.Add(uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Add(uri)
	}
	return nil
}

// AddID does the same as the AddID method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) AddID(uri string, pos int) (int, error) {
	id, err := r.client.AddID(uri, pos)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return id, err
		}
		return r.client.AddID(uri, pos)
	}
	return id, nil
}

// BeginCommandList does the same as the BeginCommandList method of
// Client.
func (r *ReconnectingClient) BeginCommandList() *CommandList {
	return r.client.BeginCommandList()
}

// Clear does the same as the Clear method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Clear() error {
	err := r.client.Clear()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Clear()
	}
	return nil
}

// Close does the same as the Close method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Close() error {
	err := r.client.Close()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Close()
	}
	return nil
}

// CurrentSong does the same as the CurrentSong method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) CurrentSong() (Attrs, error) {
	attrs, err := r.client.CurrentSong()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.CurrentSong()
	}
	return attrs, nil
}

// Delete does the same as the Delete method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Delete(start, end int) error {
	err := r.client.Delete(start, end)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Delete(start, end)
	}
	return nil
}

// DeleteID does the same as the DeleteID method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) DeleteID(id int) error {
	err := r.client.DeleteID(id)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.DeleteID(id)
	}
	return nil
}

// DisableOutput does the same as the DisableOutput method of Client,
// but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) DisableOutput(id int) error {
	err := r.client.DisableOutput(id)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.DisableOutput(id)
	}
	return nil
}

// EnableOutput does the same as the EnableOutput method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) EnableOutput(id int) error {
	err := r.client.EnableOutput(id)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.EnableOutput(id)
	}
	return nil
}

// Find does the same as the Find method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Find(uri string) ([]Attrs, error) {
	attrs, err := r.client.Find(uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.Find(uri)
	}
	return attrs, nil
}

// GetFiles does the same as the GetFiles method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) GetFiles() ([]string, error) {
	files, err := r.client.GetFiles()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return files, err
		}
		return r.client.GetFiles()
	}
	return files, nil
}

// List does the same as the List method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) List(uri string) ([]string, error) {
	list, err := r.client.List(uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return list, err
		}
		return r.client.List(uri)
	}
	return list, nil
}

// ListAllInfo does the same as the ListAllInfo method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) ListAllInfo(uri string) ([]Attrs, error) {
	attrs, err := r.client.ListAllInfo(uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.ListAllInfo(uri)
	}
	return attrs, nil
}

// ListInfo does the same as the ListInfo method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) ListInfo(uri string) ([]Attrs, error) {
	attrs, err := r.client.ListInfo(uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.ListInfo(uri)
	}
	return attrs, nil
}

// ListOutputs does the same as the ListOutputs method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) ListOutputs() ([]Attrs, error) {
	attrs, err := r.client.ListOutputs()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.ListOutputs()
	}
	return attrs, nil
}

// ListPlaylists does the same as the ListPlaylists method of Client,
// but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) ListPlaylists() ([]Attrs, error) {
	attrs, err := r.client.ListPlaylists()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.ListPlaylists()
	}
	return attrs, nil
}

// Move does the same as the Move method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Move(start, end, position int) error {
	err := r.client.Move(start, end, position)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Move(start, end, position)
	}
	return nil
}

// MoveID does the same as the MoveID method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) MoveID(songid, position int) error {
	err := r.client.MoveID(songid, position)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.MoveID(songid, position)
	}
	return nil
}

// Next does the same as the Next method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Next() error {
	err := r.client.Next()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Next()
	}
	return nil
}

// Pause does the same as the Pause method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Pause(pause bool) error {
	err := r.client.Pause(pause)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Pause(pause)
	}
	return nil
}

// Ping does the same as the Ping method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Ping() error {
	err := r.client.Ping()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Ping()
	}
	return nil
}

// Play does the same as the Play method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Play(pos int) error {
	err := r.client.Play(pos)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Play(pos)
	}
	return nil
}

// PlayID does the same as the PlayID method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) PlayID(id int) error {
	err := r.client.PlayID(id)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlayID(id)
	}
	return nil
}

// PlaylistAdd does the same as the PlaylistAdd method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistAdd(name string, uri string) error {
	err := r.client.PlaylistAdd(name, uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistAdd(name, uri)
	}
	return nil
}

// PlaylistClear does the same as the PlaylistClear method of Client,
// but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistClear(name string) error {
	err := r.client.PlaylistClear(name)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistClear(name)
	}
	return nil
}

// PlaylistContents does the same as the PlaylistContents method of
// Client, but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistContents(name string) ([]Attrs, error) {
	attrs, err := r.client.PlaylistContents(name)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.PlaylistContents(name)
	}
	return attrs, nil
}

// PlaylistDelete does the same as the PlaylistDelete method of Client,
// but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistDelete(name string, pos int) error {
	err := r.client.PlaylistDelete(name, pos)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistDelete(name, pos)
	}
	return nil
}

// PlaylistInfo does the same as the PlaylistInfo method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistInfo(start, end int) ([]Attrs, error) {
	attrs, err := r.client.PlaylistInfo(start, end)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.PlaylistInfo(start, end)
	}
	return attrs, nil
}

// PlaylistLoad does the same as the PlaylistLoad method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistLoad(name string, start, end int) error {
	err := r.client.PlaylistLoad(name, start, end)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistLoad(name, start, end)
	}
	return nil
}

// PlaylistMove does the same as the PlaylistMove method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistMove(name string, id, pos int) error {
	err := r.client.PlaylistMove(name, id, pos)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistMove(name, id, pos)
	}
	return nil
}

// PlaylistRemove does the same as the PlaylistRemove method of Client,
// but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistRemove(name string) error {
	err := r.client.PlaylistRemove(name)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistRemove(name)
	}
	return nil
}

// PlaylistRename does the same as the PlaylistRename method of Client,
// but tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistRename(name, newName string) error {
	err := r.client.PlaylistRename(name, newName)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistRename(name, newName)
	}
	return nil
}

// PlaylistSave does the same as the PlaylistSave method of Client, but
// tries to reconnect once if an error occurs.
func (r *ReconnectingClient) PlaylistSave(name string) error {
	err := r.client.PlaylistSave(name)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.PlaylistSave(name)
	}
	return nil
}

// Previous does the same as the Previous method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Previous() error {
	err := r.client.Previous()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Previous()
	}
	return nil
}

// Random does the same as the Random method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Random(random bool) error {
	err := r.client.Random(random)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Random(random)
	}
	return nil
}

// Repeat does the same as the Repeat method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Repeat(repeat bool) error {
	err := r.client.Repeat(repeat)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Repeat(repeat)
	}
	return nil
}

// Seek does the same as the Seek method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Seek(pos, time int) error {
	err := r.client.Seek(pos, time)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Seek(pos, time)
	}
	return nil
}

// SeekID does the same as the SeekID method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) SeekID(id, time int) error {
	err := r.client.SeekID(id, time)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.SeekID(id, time)
	}
	return nil
}

// SetVolume does the same as the SetVolume method of Client, but tries
// to reconnect once if an error occurs.
func (r *ReconnectingClient) SetVolume(volume int) error {
	err := r.client.SetVolume(volume)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.SetVolume(volume)
	}
	return nil
}

// Shuffle does the same as the Shuffle method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Shuffle(start, end int) error {
	err := r.client.Shuffle(start, end)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Shuffle(start, end)
	}
	return nil
}

// Stats does the same as the Stats method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Stats() (Attrs, error) {
	attrs, err := r.client.Stats()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.Stats()
	}
	return attrs, nil
}

// Status does the same as the Status method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Status() (Attrs, error) {
	attrs, err := r.client.Status()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return attrs, err
		}
		return r.client.Status()
	}
	return attrs, nil
}

// Stop does the same as the Stop method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Stop() error {
	err := r.client.Stop()
	if err != nil {
		if err = r.reconnect(); err != nil {
			return err
		}
		return r.client.Stop()
	}
	return nil
}

// Update does the same as the Update method of Client, but tries to
// reconnect once if an error occurs.
func (r *ReconnectingClient) Update(uri string) (int, error) {
	jobID, err := r.client.Update(uri)
	if err != nil {
		if err = r.reconnect(); err != nil {
			return jobID, err
		}
		return r.client.Update(uri)
	}
	return jobID, nil
}

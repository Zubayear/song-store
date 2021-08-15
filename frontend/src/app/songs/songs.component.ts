import { Component, OnInit } from '@angular/core';
import { ISong } from './song';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.css'],
})
export class SongsComponent implements OnInit {
  pageTitle = 'Song List';
  imageWidth = 50;
  imageMargin = 20;
  showImage = false;
  private _filterString = '';

  filteredSongs: ISong[] = [];

  public get filterString(): string {
    return this._filterString;
  }

  public set filterString(v: string) {
    this._filterString = v;
    this.filteredSongs = this.songFilter(v);
  }

  songs: ISong[] = [
    {
      songName: 'Humble',
      songDuration: 2.42,
      songHits: 12345678,
    },
    {
      songName: 'Middle Child',
      songDuration: 2.42,
      songHits: 12345678,
    },
    {
      songName: 'Sing about me',
      songDuration: 2.42,
      songHits: 12345678,
    },
  ];

  songFilter(filterString: string): ISong[] {
    filterString = filterString.toLocaleLowerCase();
    return this.songs.filter((song: ISong) =>
      song.songName.toLocaleLowerCase().includes(filterString)
    );
  }

  constructor() {}

  ngOnInit(): void {
    this.filterString = 'Hum';
  }
}
